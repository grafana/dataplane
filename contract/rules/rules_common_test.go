package rules_test

import (
	"testing"

	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestCommonRules(t *testing.T) {
	t.Run("CheckAtLeastOneFrame", func(t *testing.T) {
		t.Run("at_least_one_frame", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1")}
			require.Equal(t, true, rules.CheckAtLeastOneFrame.IsCompatible(frames))
		})

		t.Run("no_frames", func(t *testing.T) {
			frames := data.Frames{}
			require.Equal(t, false, rules.CheckAtLeastOneFrame.IsCompatible(frames))
		})

		t.Run("nil_frames", func(t *testing.T) {
			frames := data.Frames{nil, data.NewFrame("frame1")}
			require.Equal(t, true, rules.CheckAtLeastOneFrame.IsCompatible(frames))
		})
	})

	t.Run("CheckFramesLengthShouldBeOne", func(t *testing.T) {
		t.Run("single_frame", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1")}
			require.Equal(t, true, rules.CheckFramesLengthShouldBeOne.IsCompatible(frames))
		})

		t.Run("multiple_frames", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1"), data.NewFrame("frame2")}
			require.Equal(t, false, rules.CheckFramesLengthShouldBeOne.IsCompatible(frames))
		})
	})

	t.Run("CheckNoNilFrame", func(t *testing.T) {
		t.Run("no_nil_frames", func(t *testing.T) {
			frames := data.Frames{data.NewFrame("frame1")}
			require.Equal(t, true, rules.CheckNoNilFrame.IsCompatible(frames))
		})

		t.Run("contains_nil_frame", func(t *testing.T) {
			frames := data.Frames{nil, data.NewFrame("frame1")}
			require.Equal(t, false, rules.CheckNoNilFrame.IsCompatible(frames))
		})
	})

}
