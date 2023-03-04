package reader_test

import (
	"fmt"
	"testing"

	"github.com/grafana/dataplane/sdata/numeric"
	"github.com/grafana/dataplane/sdata/reader"
	"github.com/grafana/dataplane/sdata/timeseries"
	"github.com/grafana/dataplane/testdata/contract"
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

func TestCanReadTestData(t *testing.T) {
	e, err := contract.GetExamples()
	require.NoError(t, err)
	examples := e.GetAllAsList()
	require.Greater(t, len(examples), 1)
	for _, example := range examples {
		t.Run(example.GetInfo().Path, func(t *testing.T) {
			kind, err := reader.CanReadBasedOnMeta(example.Frames())
			require.NoError(t, err)
			require.Equal(t, example.GetInfo().Type.Kind(), kind)

			if kind == data.KindNumeric {
				nr, err := numeric.CollectionReaderFromFrames(example.Frames())
				require.NoError(t, err)

				c, err := nr.GetCollection(false)
				require.NoError(t, err)
				require.NoError(t, c.Warning)

				require.Len(t, c.Refs, int(example.GetInfo().ItemCount))
			}
		})
	}
}
