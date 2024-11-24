package rules

import (
	"time"

	"github.com/grafana/dataplane/contract/field"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// CheckEachFrameHaveAtLeastOneNonNullableTimeField ensures the following
//
// * at least one frame exist
// * each frame has at least one non-nullable time field.
var CheckEachFrameHaveAtLeastOneNonNullableTimeField = Rule{
	Name: "each frame should have at least one time non nullable field",
	IsCompatible: func(frames data.Frames) bool {
		for _, frame := range frames {
			found := false
			for _, fd := range frame.Fields {
				if fd.Type() == data.FieldTypeTime {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return len(frames) > 0 && true
	},
}

// CheckFirstTimeFieldInEachFrameShouldBeNonNil ensures the following
//
// * the first time field in each frame is non-nil
var CheckFirstTimeFieldInEachFrameShouldBeNonNil = Rule{
	Name: "first time field in each frame should be non nil",
	IsCompatible: func(frames data.Frames) bool {
		for _, frame := range frames {
			for idx, fd := range frame.Fields {
				if fd.Type() == data.FieldTypeTime {
					return true
				}
				if fd.Type() == data.FieldTypeNullableTime {
					if _, ok := fd.ConcreteAt(idx); !ok {
						return false
					}
					return true
				}
			}
		}
		return true
	},
}

// CheckFirstTimeFieldInEachFrameShouldBeSortedAscending ensures the following
//
// * first time field in each frame is sorted in ascending order.
var CheckFirstTimeFieldInEachFrameShouldBeSortedAscending = Rule{
	Name: "first time field in each frame should be sorted ascending",
	IsCompatible: func(frames data.Frames) bool {
		for _, frame := range frames {
			for _, fd := range frame.Fields {
				if fd.Type() == data.FieldTypeTime {
					for i := 1; i < fd.Len(); i++ {
						if fd.At(i - 1).(time.Time).After(fd.At(i).(time.Time)) {
							return false
						}
					}
					break
				}
			}
		}
		return true
	},
}

// CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField ensures the following
//
// * Each frame has at least one numeric or boolean field.
var CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField = Rule{
	Name: "each frame should have at least one numeric or boolean field",
	IsCompatible: func(frames data.Frames) bool {
		for _, frame := range frames {
			for _, fd := range frame.Fields {
				if field.IsNumericField(fd) || field.IsBoolField(fd) {
					return true
				}
			}
		}
		return false
	},
}
