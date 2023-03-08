package reader_test

import (
	"fmt"
	"testing"

	"github.com/grafana/dataplane/examples"
	"github.com/grafana/dataplane/sdata/numeric"
	"github.com/grafana/dataplane/sdata/reader"
	"github.com/grafana/dataplane/sdata/timeseries"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestCanReadBasedOnMeta(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		tsWideNoData, err := timeseries.NewWideFrame("A", timeseries.WideFrameVersionLatest)
		require.NoError(t, err)

		kind, err := reader.CanReadBasedOnMeta(tsWideNoData.Frames())
		require.NoError(t, err)
		require.Equal(t, data.KindTimeSeries, kind)

		nopeFrames := data.Frames{data.NewFrame("")}
		_, err = reader.CanReadBasedOnMeta(nopeFrames)
		require.Error(t, err)
	})

	t.Run("read something", func(t *testing.T) {
		n, err := numeric.NewWideFrame("A", numeric.WideFrameVersionLatest)
		require.NoError(t, err)

		err = n.AddMetric("cpu", data.Labels{"host": "king_sloth"}, 5.3)
		require.NoError(t, err)
		err = n.AddMetric("cpu", data.Labels{"host": "queen_sloth"}, 5.2)
		require.NoError(t, err)

		frames := n.Frames()

		kind, err := reader.CanReadBasedOnMeta(frames)
		require.NoError(t, err)

		if kind == data.KindNumeric {
			nr, err := numeric.CollectionReaderFromFrames(frames)
			require.NoError(t, err)

			c, err := nr.GetCollection(false)
			require.NoError(t, err)
			require.NoError(t, c.Warning)

			require.Len(t, c.Refs, 2)
			for _, ref := range c.Refs {
				_ = ref
				val, empty, err := ref.NullableFloat64Value()
				require.NoError(t, err)
				require.Equal(t, false, empty)
				fmt.Printf("%v %v %v\n", ref.GetMetricName(), ref.GetLabels().String(), *val)
			}
		}
	})
}

func TestExamplesCollections(t *testing.T) {
	examples, err := examples.GetExamples()
	require.NoError(t, err)

	collections := examples.Collections()
	_ = collections
}

func TestCanReadTestData(t *testing.T) {
	es, err := examples.GetExamples()
	require.NoError(t, err)

	numericV01Examples, err := es.Filter(examples.FilterOptions{
		Kind:    data.KindNumeric,
		Version: data.FrameTypeVersion{0, 1},
	})
	require.NoError(t, err)

	for _, collection := range numericV01Examples.Collections() {
		for _, example := range collection.ExampleSlice() {
			t.Run(example.Info().ID, func(t *testing.T) {
				kind, err := reader.CanReadBasedOnMeta(example.Frames("A"))
				require.NoError(t, err)
				require.Equal(t, example.Info().Type.Kind(), kind)
				require.Equal(t, data.KindNumeric, kind)

				nr, err := numeric.CollectionReaderFromFrames(example.Frames("A"))
				require.NoError(t, err)

				c, err := nr.GetCollection(false)
				require.NoError(t, err)
				require.NoError(t, c.Warning)

				require.Len(t, c.Refs, example.Info().ItemCount)
			})
		}
	}

	timeseriesV01Examples, err := es.Filter(examples.FilterOptions{
		Kind:    data.KindTimeSeries,
		Version: data.FrameTypeVersion{0, 1},
	})
	require.NoError(t, err)

	for _, collection := range timeseriesV01Examples.Collections() {
		for _, example := range collection.ExampleSlice() {
			t.Run(example.Info().ID, func(t *testing.T) {
				kind, err := reader.CanReadBasedOnMeta(example.Frames("A"))
				require.NoError(t, err)
				require.Equal(t, example.Info().Type.Kind(), kind)
				require.Equal(t, data.KindTimeSeries, kind)

				nr, err := timeseries.CollectionReaderFromFrames(example.Frames("A"))
				require.NoError(t, err)

				c, err := nr.GetCollection(false)
				require.NoError(t, err)
				require.NoError(t, c.Warning)

				require.Len(t, c.Refs, example.Info().ItemCount)
			})
		}
	}
}
