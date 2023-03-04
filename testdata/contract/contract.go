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
)

// The files are embedded so paths are consistent and so there
// are not issues with relative paths when using this in other's
// libraries tests.
//
//go:embed numeric/*
var content embed.FS

// ExampleInfo is additional info about the example.
// It is a mix of info that is populated from the directories
// that contain the json files and from the Meta.Custom["exampleInfo"] of the
// first frame in each file.
type ExampleInfo struct {
	Summary   string `json:"summary"`
	ItemCount int    `json:"itemCount"`

	CollectionVersion int `json:"collectionVersion"`

	// Note: Consider adding Remainder count after seeing if remainder frame/field is separate or not.

	// This following fields are populated from areas outside the Meta.Custom["exampleInfo"] (either the frame, or containing directories)
	Type       data.FrameType        `json:"-"`
	Version    data.FrameTypeVersion `json:"-"`
	Path       string                `json:"-"`
	Collection string                `json:"-"`
}

type Example struct {
	info   ExampleInfo
	frames data.Frames
}

// Info returns the ExampleInfo from an example.
func (e *Example) Info() ExampleInfo {
	return e.info
}

// Frames returns the example's data.Frames ([]*data.Frames) with each
// frame's RefID property set to refID.
// The frames returned may be modified without changing the example frames.
func (e *Example) Frames(refID string) data.Frames {
	// Reread to avoid mutation issues.
	b, err := fs.ReadFile(content, e.info.Path)
	if err != nil {
		// panic since repeat of GetExamples() which is against embed
		// so should not fail
		panic(err)
	}

	var frames data.Frames
	err = json.Unmarshal(b, &frames)
	if err != nil {
		panic(err)
	}
	if refID != "" { // all examples having default refID in frames of "" is tested for
		for _, frame := range frames {
			frame.RefID = refID
		}
	}
	return frames
}

// GetExamples returns all Examples provided by this library.
func GetExamples() (Examples, error) {
	e := Examples{}
	err := fs.WalkDir(content, "numeric", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			var frames data.Frames
			b, err := fs.ReadFile(content, path)
			if err != nil {
				return err
			}

			err = json.Unmarshal(b, &frames)
			if err != nil {
				return err
			}

			parts := strings.Split(path, string(os.PathSeparator))
			if len(parts) < 5 {
				return fmt.Errorf("unexpected test/example file path length, want at least 4 but got %v for %q", len(parts), path)
			}
			collection := parts[len(parts)-2]

			err = e.addExample(frames, collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(e) == 0 {
		return nil, fmt.Errorf("no examples found")
	}

	e.Sort(SortPathAsc)

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

// Examples is a slice of Example.
type Examples []Example

func (e *Examples) addExample(frames data.Frames, collection, path string) error {
	if e == nil {
		*e = make(Examples, 0)
	}
	example, err := newExample(frames, collection, path)
	if err != nil {
		return err
	}
	*e = append(*e, example)
	return nil
}

// FilterOptions is the argument to the Examples Filter method.
type FilterOptions struct {
	Kind              data.FrameTypeKind
	Type              data.FrameType
	Version           data.FrameTypeVersion
	Collection        string
	CollectionVersion int
}

// Filter will return a new slice of Examples filtered to
// the Examples that match any non-zero fields in FilterOptions.
func (e *Examples) Filter(f FilterOptions) (Examples, error) {
	if e == nil || len(*e) == 0 {
		return nil, fmt.Errorf("filter called empty example set")
	}

	if f.Kind != "" && f.Type != "" && f.Type.Kind() != f.Kind {
		return nil, fmt.Errorf("FrameTypeKind %q does match the FrameType %q Kind %q", f.Kind, f.Type, f.Type.Kind())
	}
	var fExamples Examples

	for _, example := range *e {
		info := example.info
		if f.Kind != "" && f.Kind != info.Type.Kind() {
			continue
		}

		if f.Type != "" && f.Type != info.Type {
			continue
		}

		if !f.Version.IsZero() && f.Version != info.Version {
			continue
		}

		if f.Collection != "" && f.Collection != info.Collection {
			continue
		}

		if f.CollectionVersion > 0 && info.CollectionVersion <= f.CollectionVersion {
			continue
		}

		fExamples = append(fExamples, example)
	}

	if len(fExamples) == 0 {
		return nil, fmt.Errorf("no examples after filtering")
	}

	return fExamples, nil
}
