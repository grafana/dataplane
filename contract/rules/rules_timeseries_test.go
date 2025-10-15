package rules_test

import (
	"testing"
	"time"

	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestRules(t *testing.T) {
	t.Run("CheckEachFrameHaveAtLeastOneNonNullableTimeField", func(t *testing.T) {
		t.Run("non_nullable_time_field_exists", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()}))}
			require.Equal(t, true, rules.CheckEachFrameHaveAtLeastOneNonNullableTimeField.IsCompatible(frames))
		})

		t.Run("no_non_nullable_time_field", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("values", nil, []*time.Time{nil}))}
			require.Equal(t, false, rules.CheckEachFrameHaveAtLeastOneNonNullableTimeField.IsCompatible(frames))
		})
	})

	t.Run("CheckFirstTimeFieldInEachFrameShouldBeSortedAscending", func(t *testing.T) {
		t.Run("time_field_sorted_ascending", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)}))}
			require.Equal(t, true, rules.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.IsCompatible(frames))
		})

		t.Run("time_field_not_sorted", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now().Add(time.Minute), time.Now()}))}
			require.Equal(t, false, rules.CheckFirstTimeFieldInEachFrameShouldBeSortedAscending.IsCompatible(frames))
		})
	})

	t.Run("CheckFirstTimeFieldInEachFrameShouldBeNonNil", func(t *testing.T) {
		t.Run("time_field_non_nil", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now(), time.Now().Add(time.Minute)}))}
			require.Equal(t, true, rules.CheckFirstTimeFieldInEachFrameShouldBeNonNil.IsCompatible(frames))
		})

		t.Run("time_field_has_nil_values", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []*time.Time{nil, nil}))}
			require.Equal(t, false, rules.CheckFirstTimeFieldInEachFrameShouldBeNonNil.IsCompatible(frames))
		})
	})

	t.Run("CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField", func(t *testing.T) {
		t.Run("numeric_field_exists", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("value", nil, []float64{1.0}))}
			require.Equal(t, true, rules.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.IsCompatible(frames))
		})

		t.Run("no_numeric_or_boolean_field", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []time.Time{time.Now()}))}
			require.Equal(t, false, rules.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.IsCompatible(frames))
		})

		t.Run("nullable_numeric_field_exists", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("value", nil, []*float64{nil, new(float64)}))}
			require.Equal(t, true, rules.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.IsCompatible(frames))
		})

		t.Run("nullable_fields_no_numeric_or_boolean", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1", data.NewField("time", nil, []*time.Time{nil, nil}))}
			require.Equal(t, false, rules.CheckEachFrameShouldHaveAtLeastOneNumericOrBooleanField.IsCompatible(frames))
		})
	})
}
