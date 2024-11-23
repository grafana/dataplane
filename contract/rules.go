package contract

import (
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// Rule defines a rule for checking compliance.
type Rule struct {
	Name       string
	Compatible func(frames data.Frames) Compliance
}

// Rules is a collection of ComplianceRule.
type Rules []Rule

func (rules *Rules) IsCompliant(frames data.Frames) ContractErrors {
	errors := ContractErrors{}
	for _, rule := range *rules {
		if rule.Compatible(frames) == Invalid {
			errors = append(errors, ContractError{Message: rule.Name})
		}
	}
	return errors
}

// CheckAtLeastOneFrame ensures there is at least one frame.
var CheckAtLeastOneFrame = Rule{
	Name: "ensure there is at least one frame",
	Compatible: func(frames data.Frames) Compliance {
		if len(frames) > 0 {
			return Valid
		}
		return Invalid
	},
}

// CheckNoNilFrame ensures there are no nil frames.
var CheckNoNilFrame = Rule{
	Name: "ensure there no nil frames",
	Compatible: func(frame data.Frames) Compliance {
		for _, frame := range frame {
			if frame == nil {
				return Invalid
			}
		}
		return Valid
	},
}

// CheckFramesLengthShouldBeOne ensures the frames length is one.
var CheckFramesLengthShouldBeOne = Rule{
	Name: "frames length should be one",
	Compatible: func(frames data.Frames) Compliance {
		if len(frames) == 1 {
			return Valid
		}
		return Invalid
	},
}

// CheckEachFrameHaveAtLeastOneNonNullableTimeField ensures each frame has at least one non-nullable time field.
var CheckEachFrameHaveAtLeastOneNonNullableTimeField = Rule{
	Name: "each frame should have at least one time non nullable field",
	Compatible: func(frames data.Frames) Compliance {
		for _, frame := range frames {
			found := false
			for _, field := range frame.Fields {
				if field.Type() == data.FieldTypeTime {
					found = true
					break
				}
			}
			if !found {
				return Invalid
			}
		}
		return Valid
	},
}

// CheckFirstTimeFieldInEachFrameShouldBeNonNil ensures the first time field in each frame is non-nil.
var CheckFirstTimeFieldInEachFrameShouldBeNonNil = Rule{
	Name: "first time field in each frame should be non nil",
	Compatible: func(frames data.Frames) Compliance {
		for _, frame := range frames {
			for idx, field := range frame.Fields {
				if field.Type() == data.FieldTypeTime {
					return Valid
				}
				if field.Type() == data.FieldTypeNullableTime {
					if _, ok := field.ConcreteAt(idx); !ok {
						return Invalid
					}
					return Valid
				}
			}
		}
		return Valid
	},
}

// CheckFirstTimeFieldInEachFrameShouldBeSortedAscending ensures the first time field in each frame is sorted in ascending order.
var CheckFirstTimeFieldInEachFrameShouldBeSortedAscending = Rule{
	Name: "first time field in each frame should be sorted ascending",
	Compatible: func(frames data.Frames) Compliance {
		for _, frame := range frames {
			for _, field := range frame.Fields {
				if field.Type() == data.FieldTypeTime {
					for i := 1; i < field.Len(); i++ {
						if field.At(i - 1).(time.Time).After(field.At(i).(time.Time)) {
							return Invalid
						}
					}
					break
				}
			}
		}
		return Valid
	},
}

// CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField ensures each frame has at least one numeric or boolean field.
var CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField = Rule{
	Name: "each frame should have at least one numeric or boolean field",
	Compatible: func(frames data.Frames) Compliance {
		for _, frame := range frames {
			for _, field := range frame.Fields {
				if IsNumericField(field) || IsBoolField(field) {
					return Valid
				}
			}
		}
		return Invalid
	},
}
