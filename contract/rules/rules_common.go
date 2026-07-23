package rules

import (
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// CheckAtLeastOneFrame ensures the following
//
// * at least one frame exist
var CheckAtLeastOneFrame = Rule{
	Name: "ensure there is at least one frame",
	IsCompatible: func(frames data.Frames) bool {
		return len(frames) > 0
	},
}

// CheckFramesLengthShouldBeOne ensures the following
//
// * frames length is one
var CheckFramesLengthShouldBeOne = Rule{
	Name: "frames length should be one",
	IsCompatible: func(frames data.Frames) bool {
		return len(frames) == 1
	},
}

// CheckNoNilFrame ensures the following
//
// * at least one frame exist
// * no nil frames exist
var CheckNoNilFrame = Rule{
	Name: "ensure there no nil frames",
	IsCompatible: func(frames data.Frames) bool {
		for _, frame := range frames {
			if frame == nil {
				return false
			}
		}
		return len(frames) > 0 && true
	},
}
