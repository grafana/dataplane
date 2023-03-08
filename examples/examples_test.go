package examples_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/grafana/dataplane/examples"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestGetExamples(t *testing.T) {
	_, err := examples.GetExamples()
	require.NoError(t, err)
}

func TestExamplesSort(t *testing.T) {
	es, err := examples.GetExamples()
	require.NoError(t, err)

	numericExamples, err := es.Filter(examples.FilterOptions{Kind: data.KindNumeric})
	require.NoError(t, err)

	numericExamples.Sort(examples.SortFrameTypeAsc)
	require.Equal(t, data.FrameTypeNumericLong, numericExamples.AsSlice()[0].Info().Type)

	numericExamples.Sort(examples.SortFrameTypeDesc)
	require.Equal(t, data.FrameTypeNumericWide, numericExamples.AsSlice()[0].Info().Type)
}

func TestValidExamples(t *testing.T) {
	examples, err := examples.GetExamples()

	require.NoError(t, err)

	t.Run("all sumaries must end in period", func(t *testing.T) {
		for _, e := range examples.AsSlice() {
			info := e.Info()
			require.True(t, strings.HasSuffix(info.Summary, "."), fmt.Sprintf("Summary: %q Path: %q", info.Summary, info.Path))
		}
	})
	t.Run("all sumaries must have collectionVersion >= 1", func(t *testing.T) {
		for _, e := range examples.AsSlice() {
			info := e.Info()
			require.GreaterOrEqual(t, e.Info().CollectionVersion, 1, info.Path)
		}
	})

	t.Run("all frames have zero value (empty string) refID", func(t *testing.T) {
		for _, e := range examples.AsSlice() {
			for _, frame := range e.Frames("") {
				require.Empty(t, frame.RefID)
			}
		}
	})
}

func TestExampleFramesMutation(t *testing.T) {
	examples, err := examples.GetExamples()
	s := examples.AsSlice()
	require.NoError(t, err)
	ft := s[0].Frames("")[0].Meta.Type
	require.NotEmpty(t, ft)

	frame := s[0].Frames("")[0]
	frame.Meta.Type = "sloth"
	newFT := s[0].Frames("")[0].Meta.Type
	require.Equal(t, ft, newFT)
}

func TestExamplesFilter(t *testing.T) {
	e, err := examples.GetExamples()
	require.NoError(t, err)

	t.Run("frametype and kind must match if not zero value", func(t *testing.T) {
		_, err := e.Filter(examples.FilterOptions{Type: data.FrameTypeTimeSeriesLong, Kind: data.KindNumeric})
		require.Error(t, err)
	})

	t.Run("version filter will error for no results for version with no match", func(t *testing.T) {
		_, err := e.Filter(examples.FilterOptions{Version: data.FrameTypeVersion{124, 1233}})
		require.Error(t, err)
	})

	t.Run("can filter to numeric long", func(t *testing.T) {
		numLongExamples, err := e.Filter(examples.FilterOptions{Type: data.FrameTypeNumericLong})

		require.NoError(t, err)
		for _, e := range numLongExamples.AsSlice() {
			require.Equal(t, data.FrameTypeNumericLong, e.Info().Type)
		}
	})
}
