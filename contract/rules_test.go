package contract_test

import (
	"testing"
	"time"

	c "github.com/grafana/dataplane/contract"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestChecks(t *testing.T) {
	t.Run("Checks", func(t *testing.T) {
		t.Run("Check at least one frame exist", func(t *testing.T) {
			t.Run("at_least_one_frame", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1")}
				require.Equal(t, c.Valid, c.CheckAtLeastOneFrame.Compatible(frames))
			})

			t.Run("no_frames", func(t *testing.T) {
				frames := data.Frames{}
				require.Equal(t, c.Invalid, c.CheckAtLeastOneFrame.Compatible(frames))
			})

			t.Run("nil_frames", func(t *testing.T) {
				frames := data.Frames{nil, data.NewFrame("frame1")}
				require.Equal(t, c.Valid, c.CheckAtLeastOneFrame.Compatible(frames))
			})
		})

		t.Run("Check no nil frame exist", func(t *testing.T) {
			t.Run("no_nil_frames", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1")}
				require.Equal(t, c.Valid, c.CheckNoNilFrame.Compatible(frames))
			})

			t.Run("contains_nil_frame", func(t *testing.T) {
				frames := data.Frames{nil, data.NewFrame("frame1")}
				require.Equal(t, c.Invalid, c.CheckNoNilFrame.Compatible(frames))
			})
		})

		t.Run("Check frame length is 1", func(t *testing.T) {
			t.Run("single_frame", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1")}
				require.Equal(t, c.Valid, c.CheckFramesLengthShouldBeOne.Compatible(frames))
			})

			t.Run("multiple_frames", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1"), data.NewFrame("frame2")}
				require.Equal(t, c.Invalid, c.CheckFramesLengthShouldBeOne.Compatible(frames))
			})
		})

		t.Run("Check each frame have at least one non-nullable tine field", func(t *testing.T) {
			t.Run("non_nullable_time_field_exists", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()}))}
				require.Equal(t, c.Valid, c.CheckEachFrameHaveAtLeastOneNonNullableTimeField.Compatible(frames))
			})

			t.Run("no_non_nullable_time_field", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("values", nil, []*time.Time{nil}))}
				require.Equal(t, c.Invalid, c.CheckEachFrameHaveAtLeastOneNonNullableTimeField.Compatible(frames))
			})

			t.Run("multiple_frames_with_non_nullable_time_field", func(t *testing.T) {
				frames := data.Frames{
					data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()})),
					data.NewFrame("frame2", data.NewField("time", nil, []time.Time{time.Now()})),
				}
				require.Equal(t, c.Valid, c.CheckEachFrameHaveAtLeastOneNonNullableTimeField.Compatible(frames))
			})

			t.Run("multiple_frames_with_non_nullable_time_field", func(t *testing.T) {
				frames := data.Frames{
					data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()})),
					data.NewFrame("frame2", data.NewField("time", nil, []*time.Time{nil})),
				}
				require.Equal(t, c.Invalid, c.CheckEachFrameHaveAtLeastOneNonNullableTimeField.Compatible(frames))
			})
		})

		t.Run("Check first time field sorted by ascending", func(t *testing.T) {
			t.Run("time_field_sorted_ascending", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)}))}
				require.Equal(t, c.Valid, c.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.Compatible(frames))
			})

			t.Run("time_field_not_sorted", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now().Add(time.Minute), time.Now()}))}
				require.Equal(t, c.Invalid, c.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.Compatible(frames))
			})

			t.Run("multiple_frames_sorted_ascending", func(t *testing.T) {
				frames := data.Frames{
					data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)})),
					data.NewFrame("frame2", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)})),
				}
				require.Equal(t, c.Valid, c.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.Compatible(frames))
			})

			t.Run("multiple_frames_sorted_ascending", func(t *testing.T) {
				frames := data.Frames{
					data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)})),
					data.NewFrame("frame2", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Hour), time.Now().Add(time.Minute)})),
				}
				require.Equal(t, c.Invalid, c.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.Compatible(frames))
			})
		})

		t.Run("Check first time field in not non-nullable", func(t *testing.T) {
			t.Run("time_field_non_nil", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)}))}
				require.Equal(t, c.Valid, c.CheckFirstTimeFieldInEachFrameShouldBeNonNil.Compatible(frames))
			})

			t.Run("time_field_has_nil_values", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []*time.Time{nil, nil}))}
				require.Equal(t, c.Invalid, c.CheckFirstTimeFieldInEachFrameShouldBeNonNil.Compatible(frames))
			})

			t.Run("multiple_frames_non_nil", func(t *testing.T) {
				frames := data.Frames{
					data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)})),
					data.NewFrame("frame2", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)})),
				}
				require.Equal(t, c.Valid, c.CheckFirstTimeFieldInEachFrameShouldBeNonNil.Compatible(frames))
			})
		})

		t.Run("Check each frame at least one numeric or boolean field", func(t *testing.T) {
			t.Run("numeric_field_exists", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("value", nil, []float64{1.0}))}
				require.Equal(t, c.Valid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
			})

			t.Run("no_numeric_or_boolean_field", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()}))}
				require.Equal(t, c.Invalid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
			})

			t.Run("nullable_numeric_field_exists", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("value", nil, []*float64{nil, new(float64)}))}
				require.Equal(t, c.Valid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
			})

			t.Run("nullable_fields_no_numeric_or_boolean", func(t *testing.T) {
				frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []*time.Time{nil, nil}))}
				require.Equal(t, c.Invalid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
			})

			// Additional test for multiple frames
			t.Run("multiple_frames_numeric_field_exists", func(t *testing.T) {
				frames := data.Frames{
					data.NewFrame("frame1", data.NewField("value", nil, []float64{1.0})),
					data.NewFrame("frame2", data.NewField("value", nil, []float64{2.0})),
				}
				require.Equal(t, c.Valid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
			})
		})

	})
	t.Run("CheckAtLeastOneFrame", func(t *testing.T) {
		t.Run("at_least_one_frame", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1")}
			require.Equal(t, c.Valid, c.CheckAtLeastOneFrame.Compatible(frames))
		})

		t.Run("no_frames", func(t *testing.T) {
			frames := data.Frames{}
			require.Equal(t, c.Invalid, c.CheckAtLeastOneFrame.Compatible(frames))
		})

		t.Run("nil_frames", func(t *testing.T) {
			frames := data.Frames{nil, data.NewFrame("frame1")}
			require.Equal(t, c.Valid, c.CheckAtLeastOneFrame.Compatible(frames))
		})
	})

	t.Run("CheckNoNilFrame", func(t *testing.T) {
		t.Run("no_nil_frames", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1")}
			require.Equal(t, c.Valid, c.CheckNoNilFrame.Compatible(frames))
		})

		t.Run("contains_nil_frame", func(t *testing.T) {
			frames := data.Frames{nil, data.NewFrame("frame1")}
			require.Equal(t, c.Invalid, c.CheckNoNilFrame.Compatible(frames))
		})
	})

	t.Run("CheckFramesLengthShouldBeOne", func(t *testing.T) {
		t.Run("single_frame", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1")}
			require.Equal(t, c.Valid, c.CheckFramesLengthShouldBeOne.Compatible(frames))
		})

		t.Run("multiple_frames", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1"), data.NewFrame("frame2")}
			require.Equal(t, c.Invalid, c.CheckFramesLengthShouldBeOne.Compatible(frames))
		})
	})

	t.Run("CheckEachFrameShouldHaveAtLeastOneNonNullableTimeField", func(t *testing.T) {
		t.Run("non_nullable_time_field_exists", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()}))}
			require.Equal(t, c.Valid, c.CheckEachFrameHaveAtLeastOneNonNullableTimeField.Compatible(frames))
		})

		t.Run("no_non_nullable_time_field", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("values", nil, []*time.Time{nil}))}
			require.Equal(t, c.Invalid, c.CheckEachFrameHaveAtLeastOneNonNullableTimeField.Compatible(frames))
		})
	})

	t.Run("CheckFirstTimeFieldInEachFrameShouldBeSortedAscending", func(t *testing.T) {
		t.Run("time_field_sorted_ascending", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)}))}
			require.Equal(t, c.Valid, c.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.Compatible(frames))
		})

		t.Run("time_field_not_sorted", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now().Add(time.Minute), time.Now()}))}
			require.Equal(t, c.Invalid, c.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.Compatible(frames))
		})
	})

	t.Run("CheckFirstTimeFieldInEachFrameShouldBeNonNil", func(t *testing.T) {
		t.Run("time_field_non_nil", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)}))}
			require.Equal(t, c.Valid, c.CheckFirstTimeFieldInEachFrameShouldBeNonNil.Compatible(frames))
		})

		t.Run("time_field_has_nil_values", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []*time.Time{nil, nil}))}
			require.Equal(t, c.Invalid, c.CheckFirstTimeFieldInEachFrameShouldBeNonNil.Compatible(frames))
		})
	})

	t.Run("CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField", func(t *testing.T) {
		t.Run("numeric_field_exists", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("value", nil, []float64{1.0}))}
			require.Equal(t, c.Valid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
		})

		t.Run("no_numeric_or_boolean_field", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()}))}
			require.Equal(t, c.Invalid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
		})

		t.Run("nullable_numeric_field_exists", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("value", nil, []*float64{nil, new(float64)}))}
			require.Equal(t, c.Valid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
		})

		t.Run("nullable_fields_no_numeric_or_boolean", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []*time.Time{nil, nil}))}
			require.Equal(t, c.Invalid, c.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.Compatible(frames))
		})
	})
}
