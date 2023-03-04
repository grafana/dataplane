// Package contract provides example data dataplane contract
// data for testing
package contract

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	jsoniter "github.com/json-iterator/go"
)

//go:embed numeric/*
var content embed.FS

type ExampleInfo struct {
	Summary   string `json:"summary"`
	ItemCount int64  `json:"itemCount"`
	// Note: Consider adding Remainder count after seeing if remainder frame/field is separate or not.

	// Note: Consider adding some sort of "sets" and "set version"
	// this would be another (leaf) folder. So for example can have sets "basic_valid", "invalid",
	// and "extended" sets. Having a version for the set would be so that when an example is added,
	// util tests functions could log/warn instead of breaking until they opt-in to the new tests.
	// although maintainers can just get all if they wish.

	Type       data.FrameType        `json:"-"`
	Version    data.FrameTypeVersion `json:"-"`
	Path       string                `json:"-"`
	Collection string                `json:"-"`
}

type Example struct {
	info   ExampleInfo
	frames data.Frames
}

func (e Example) GetInfo() ExampleInfo {
	return e.info
}

func (e Example) Frames() data.Frames {
	return e.frames
}

func GetExamples() (Examples, error) {
	e := Examples{}
	err := fs.WalkDir(content, "numeric", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			frames := make(data.Frames, 0)
			b, err := fs.ReadFile(content, path)
			if err != nil {
				return err
			}

			err = testIterRead(&frames, b)
			if err != nil {
				return err
			}

			parts := strings.Split(path, string(os.PathSeparator))
			if len(parts) < 5 {
				return fmt.Errorf("unexpected test/example file path length, want at least 4 but got %v for %q", len(parts), path)
			}
			collection := parts[len(parts)-2]
			rawVersion := parts[len(parts)-3]
			frameType := data.FrameType(parts[len(parts)-4])

			ver, err := data.ParseFrameTypeVersion(strings.TrimPrefix(rawVersion, "v"))
			if err != nil {
				return err
			}

			err = e.addExample(frameType, ver, frames, collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return e, err
	}
	return e, nil
}

func newExample(frames data.Frames, collection, path string) (Example, error) {
	e := Example{
		frames: frames,
	}
	ei, err := exampleInfoFromFrames(frames, collection, path)
	if err != nil {
		return e, err
	}
	e.info = ei
	return e, nil
}

func exampleInfoFromFrames(frames data.Frames, collection, path string) (ExampleInfo, error) {
	info := ExampleInfo{}
	if len(frames) == 0 {
		return info, fmt.Errorf("Example (frames) is nil or zero length and must have at least one frame for path %s", path)
	}

	firstFrame := frames[0]
	if firstFrame == nil {
		return info, fmt.Errorf("nil frame should not exist for path %s", path)
	}
	if firstFrame.Meta == nil {
		return info, fmt.Errorf("first first meta is nil so missing example info for path %s", path)
	}

	if firstFrame.Meta.Custom == nil {
		return info, fmt.Errorf("custom meta data is missing so missing example info for path %s", path)
	}

	custom, ok := firstFrame.Meta.Custom.(map[string]interface{})
	if !ok {
		return info, fmt.Errorf(`custom meta data is not an object ({"string": value}) so missing example info for path %s`, path)
	}

	infoRaw, found := custom["exampleInfo"]
	if !found {
		return info, fmt.Errorf(`exampleInfo property not found is custom metadata so missing example info for path %s`, path)
	}

	b, err := json.Marshal(infoRaw)
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(b, &info)
	info.Type = firstFrame.Meta.Type
	info.Version = firstFrame.Meta.TypeVersion
	info.Collection = collection
	info.Path = path

	return info, err
}

type Examples struct {
	m map[data.FrameTypeKind]map[data.FrameType]map[data.FrameTypeVersion]map[string][]Example
}

func (e *Examples) addExample(t data.FrameType, v data.FrameTypeVersion, frames data.Frames, collection, path string) error {
	if e.m == nil {
		e.m = make(map[data.FrameTypeKind]map[data.FrameType]map[data.FrameTypeVersion]map[string][]Example)
	}

	if e.m[t.Kind()] == nil {
		e.m[t.Kind()] = make(map[data.FrameType]map[data.FrameTypeVersion]map[string][]Example)
	}

	if e.m[t.Kind()][t] == nil {
		e.m[t.Kind()][t] = make(map[data.FrameTypeVersion]map[string][]Example)
	}

	if e.m[t.Kind()][t][v] == nil {
		e.m[t.Kind()][t][v] = make(map[string][]Example)
	}

	example, err := newExample(frames, collection, path)
	if err != nil {
		return err
	}

	e.m[t.Kind()][t][v][collection] = append(e.m[t.Kind()][t][v][collection], example)
	return nil
}

func (e *Examples) GetAllAsList() []Example {
	es := []Example{}
	for kind, typeToVersion := range e.m {
		for fType, versionToCollection := range typeToVersion {
			for version, collectionToExamples := range versionToCollection {
				for collection, examples := range collectionToExamples {
					es = append(es, examples...)
					_, _, _, _ = kind, version, fType, collection
				}
			}
		}
	}
	return es
}

func testIterRead(d *data.Frames, b []byte) error {
	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, b)
	for iter.ReadArray() {
		frame := &data.Frame{}
		iter.ReadVal(frame)
		if iter.Error != nil {
			return iter.Error
		}
		*d = append(*d, frame)
	}
	return nil
}
