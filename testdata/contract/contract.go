// Package contract provides example data dataplane contract
// data for testing
package contract

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	jsoniter "github.com/json-iterator/go"
)

var testDataRelPath = "."

type ExampleInfo struct {
	Summary   string `json:"summary"`
	ItemCount int64  `json:"itemCount"`
	// Note: Consider adding Remainder count after seeing if remainder frame/field is separate or not.

	Type    data.FrameType        `json:"-"`
	Version data.FrameTypeVersion `json:"-"`
	Path    string                `json:"-"`
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
	err := filepath.Walk(testDataRelPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			frames := make(data.Frames, 0)
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			err = testIterRead(&frames, b)
			if err != nil {
				return err
			}

			parts := strings.Split(path, string(os.PathSeparator))
			if len(parts) < 4 {
				return fmt.Errorf("unexpected test/example file path length, want at least 4 but got %v for %q", len(parts), path)
			}
			name := strings.TrimSuffix(parts[len(parts)-1], ".json")
			rawVersion := parts[len(parts)-2]
			kind := data.FrameTypeKind(parts[len(parts)-3])

			ver, err := data.ParseFrameTypeVersion(strings.TrimPrefix(rawVersion, "v"))
			if err != nil {
				return err
			}

			err = e.addExample(kind, ver, name, frames, path)
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

func newExample(frames data.Frames, path string) (Example, error) {
	e := Example{
		frames: frames,
	}
	ei, err := exampleInfoFromFrames(frames, path)
	if err != nil {
		return e, err
	}
	e.info = ei
	return e, nil
}

func exampleInfoFromFrames(frames data.Frames, path string) (ExampleInfo, error) {
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
	info.Path = path

	return info, err
}

type Examples struct {
	m map[data.FrameTypeKind]map[data.FrameTypeVersion]map[string]Example
}

func (e *Examples) addExample(k data.FrameTypeKind, v data.FrameTypeVersion, name string, frames data.Frames, path string) error {
	if e.m == nil {
		e.m = make(map[data.FrameTypeKind]map[data.FrameTypeVersion]map[string]Example)
	}
	if e.m[k] == nil {
		e.m[k] = make(map[data.FrameTypeVersion]map[string]Example)
	}

	if e.m[k][v] == nil {
		e.m[k][v] = make(map[string]Example)
	}
	example, err := newExample(frames, path)
	if err != nil {
		return err
	}
	e.m[k][v][name] = example
	return nil
}

func (e *Examples) GetAllAsList() []Example {
	es := []Example{}
	for kind, versionToName := range e.m {
		for version, nameToExample := range versionToName {
			for name, example := range nameToExample {
				es = append(es, example)
				_, _, _ = kind, version, name
			}
		}
	}
	return es
}

func (s Examples) GetFrames(refID string, k data.FrameTypeKind, v data.FrameTypeVersion, scenarioID string) {

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
