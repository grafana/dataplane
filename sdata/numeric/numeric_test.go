package numeric_test

import (
	"testing"

	"github.com/grafana/dataplane/sdata/numeric"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestSimpleNumeric(t *testing.T) {
	// addMetrics uses the writer interface to add sample metrics
	addMetrics := func(c numeric.CollectionWriter) {
		err := c.AddMetric("os.cpu", data.Labels{"host": "a"}, 1.0)
		require.NoError(t, err)
		err = c.AddMetric("os.cpu", data.Labels{"host": "b"}, 2.0)
		require.NoError(t, err)
	}

	// refs should be same across the formats
	expectedRefs := []numeric.MetricRef{
		{
			ValueField: data.NewField("os.cpu", data.Labels{"host": "a"}, []float64{1}),
		},
		{
			ValueField: data.NewField("os.cpu", data.Labels{"host": "b"}, []float64{2}),
		},
	}

	t.Run("multi frame", func(t *testing.T) {
		var mFrameNC numeric.CollectionRW
		var err error
		mFrameNC, err = numeric.NewMultiFrame("A", numeric.MultiFrameVersionLatest)
		require.NoError(t, err)

		addMetrics(mFrameNC)

		mc, err := mFrameNC.GetCollection(false)
		require.NoError(t, mc.Warning)
		require.Nil(t, err)
		require.Nil(t, mc.RemainderIndices)
		require.Equal(t, expectedRefs, mc.Refs)
	})

	t.Run("wide frame", func(t *testing.T) {
		var wFrameNC numeric.CollectionRW
		var err error
		wFrameNC, err = numeric.NewWideFrame("B", numeric.WideFrameVersionLatest)
		require.NoError(t, err)

		addMetrics(wFrameNC)

		wc, err := wFrameNC.GetCollection(false)
		require.NoError(t, wc.Warning)
		require.Nil(t, err)
		require.Nil(t, wc.RemainderIndices)
		require.Equal(t, expectedRefs, wc.Refs)
	})
	t.Run("long frame", func(t *testing.T) {
		lfn, err := numeric.NewLongFrame("C", numeric.LongFrameVersionLatest)
		require.NoError(t, err)

		lfn.Fields = append(lfn.Fields, data.NewField("os.cpu", nil, []float64{1, 2}),
			data.NewField("host", nil, []string{"a", "b"}))
		var lcr numeric.CollectionReader = lfn

		lc, err := lcr.GetCollection(false)
		require.NoError(t, lc.Warning)
		require.Nil(t, err)
		require.Nil(t, lc.RemainderIndices)
		require.Equal(t, expectedRefs, lc.Refs)
	})
}

func TestNoDataFromNew(t *testing.T) {
	var multi, wide, long numeric.CollectionReader
	var err error

	multi, err = numeric.NewMultiFrame("A", numeric.MultiFrameVersionLatest)
	require.NoError(t, err)

	wide, err = numeric.NewWideFrame("B", numeric.WideFrameVersionLatest)
	require.NoError(t, err)

	long, err = numeric.NewLongFrame("C", numeric.LongFrameVersionLatest)
	require.NoError(t, err)

	noDataReqs := func(c numeric.Collection, err error) {
		require.NoError(t, err)
		require.Nil(t, c.RemainderIndices)
		require.NotNil(t, c.Refs)
		require.Len(t, c.Refs, 0)
		require.NoError(t, c.Warning)
		require.True(t, c.NoData())
	}

	viaFrames := func(r numeric.CollectionReader) {
		t.Run("should work when losing go type via Frames()", func(t *testing.T) {
			frames := r.Frames()
			r, err := numeric.CollectionReaderFromFrames(frames)
			require.NoError(t, err)

			c, err := r.GetCollection(false)
			noDataReqs(c, err)
		})
	}

	t.Run("multi", func(t *testing.T) {
		c, err := multi.GetCollection(false)
		noDataReqs(c, err)
		viaFrames(multi)
	})

	t.Run("wide", func(t *testing.T) {
		c, err := wide.GetCollection(false)
		noDataReqs(c, err)
		viaFrames(wide)
	})

	t.Run("long", func(t *testing.T) {
		c, err := long.GetCollection(false)
		noDataReqs(c, err)
		viaFrames(long)
	})
}

func TestSortNumericMetricRef(t *testing.T) {
	t.Run("sort by metric name", func(t *testing.T) {
		metricA := numeric.MetricRef{ValueField: data.NewField("metric_a", nil, []float64{})}
		metricB := numeric.MetricRef{ValueField: data.NewField("metric_b", nil, []float64{})}
		metricC := numeric.MetricRef{ValueField: data.NewField("metric_c", nil, []float64{})}
		input := []numeric.MetricRef{
			metricC,
			metricA,
			metricB,
		}
		numeric.SortNumericMetricRef(input)
		expected := []numeric.MetricRef{
			metricA,
			metricB,
			metricC,
		}

		require.Equal(t, expected, input)
	})

	t.Run("sort by labels", func(t *testing.T) {
		metricA := numeric.MetricRef{ValueField: data.NewField("metric", data.Labels{"a": "1"}, []float64{})}
		metricB := numeric.MetricRef{ValueField: data.NewField("metric", data.Labels{"b": "2"}, []float64{})}
		metricC := numeric.MetricRef{ValueField: data.NewField("metric", data.Labels{"c": "3"}, []float64{})}
		metricNilLabel := numeric.MetricRef{ValueField: data.NewField("metric", nil, []float64{})}
		input := []numeric.MetricRef{
			metricNilLabel,
			metricC,
			metricA,
			metricNilLabel,
			metricB,
		}
		numeric.SortNumericMetricRef(input)
		expected := []numeric.MetricRef{
			metricNilLabel,
			metricNilLabel,
			metricA,
			metricB,
			metricC,
		}

		require.Equal(t, expected, input)
	})
}
