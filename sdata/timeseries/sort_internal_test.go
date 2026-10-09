package timeseries

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestSortTimeSeriesMetricRefKeepsOrder(t *testing.T) {
	tests := []struct {
		name   string
		fields []*data.Field
	}{
		{name: "by metric name", fields: []*data.Field{
			data.NewField("b", data.Labels{"host": "a"}, []float64{1}),
			data.NewField("a", nil, []float64{1}),
			data.NewField("c", data.Labels{"host": "a"}, []float64{1}),
		}},
		{name: "by label string", fields: []*data.Field{
			data.NewField("value", data.Labels{"host": "b"}, []float64{1}),
			data.NewField("value", data.Labels{"host": "a", "env": "prod"}, []float64{1}),
			data.NewField("value", data.Labels{"host": "a"}, []float64{1}),
		}},
		{name: "nil, empty and set labels", fields: []*data.Field{
			data.NewField("value", data.Labels{"host": "a"}, []float64{1}),
			data.NewField("value", nil, []float64{1}),
			data.NewField("value", data.Labels{}, []float64{1}),
			data.NewField("value", nil, []float64{1}),
		}},
		{name: "missing value field", fields: []*data.Field{
			data.NewField("value", data.Labels{"host": "a"}, []float64{1}),
			nil,
			data.NewField("", nil, []float64{1}),
		}},
		{name: "different labels with the same string", fields: []*data.Field{
			data.NewField("value", data.Labels{"a": "b, c=d"}, []float64{1}),
			data.NewField("value", data.Labels{"a": "a"}, []float64{1}),
			data.NewField("value", data.Labels{"a": "b", "c": "d"}, []float64{1}),
		}},
		{name: "equal keys", fields: []*data.Field{
			data.NewField("value", data.Labels{"host": "a"}, []float64{1}),
			data.NewField("value", data.Labels{"host": "a"}, []float64{2}),
			data.NewField("value", data.Labels{"host": "a"}, []float64{3}),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireSameOrder(t, tt.fields)
		})
	}
	t.Run("random inputs", func(t *testing.T) {
		rng := rand.New(rand.NewPCG(1, 2))
		for range 500 {
			requireSameOrder(t, randomFields(rng, 1+rng.IntN(300)))
		}
	})
}

func TestSortTimeSeriesMetricRefBuildsEachKeyOnce(t *testing.T) {
	refs := toRefs(randomFields(rand.New(rand.NewPCG(1, 2)), 500))
	keyLen := 0
	keyAllocs := testing.AllocsPerRun(10, func() {
		for _, r := range refs {
			if l := r.GetLabels(); l != nil {
				keyLen += len(l.String())
			}
		}
	})
	buf := make([]MetricRef, len(refs))
	sortAllocs := testing.AllocsPerRun(10, func() {
		copy(buf, refs)
		sortTimeSeriesMetricRef(buf)
	})
	require.Positive(t, keyLen)
	require.LessOrEqual(t, sortAllocs, keyAllocs+10)
}

func BenchmarkSortTimeSeriesMetricRef(b *testing.B) {
	refs := toRefs(randomFields(rand.New(rand.NewPCG(1, 2)), 500))
	buf := make([]MetricRef, len(refs))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		copy(buf, refs)
		sortTimeSeriesMetricRef(buf)
	}
}

// sortWithStringPerComparison builds label strings inside the comparator and defines the expected order.
func sortWithStringPerComparison(refs []MetricRef) {
	sort.SliceStable(refs, func(i, j int) bool {
		iRef, jRef := refs[i], refs[j]
		if iRef.GetMetricName() < jRef.GetMetricName() {
			return true
		}
		if iRef.GetMetricName() > jRef.GetMetricName() {
			return false
		}
		if iRef.GetLabels() == nil && jRef.GetLabels() == nil {
			return true
		}
		if iRef.GetLabels() == nil && jRef.GetLabels() != nil {
			return true
		}
		if iRef.GetLabels() != nil && jRef.GetLabels() == nil {
			return false
		}
		return iRef.GetLabels().String() < jRef.GetLabels().String()
	})
}

func requireSameOrder(t *testing.T, fields []*data.Field) {
	t.Helper()
	want := toRefs(fields)
	got := slices.Clone(want)
	sortWithStringPerComparison(want)
	sortTimeSeriesMetricRef(got)
	for i := range want {
		require.Same(t, want[i].ValueField, got[i].ValueField, "index %d", i)
	}
}

func toRefs(fields []*data.Field) []MetricRef {
	refs := make([]MetricRef, len(fields))
	for i, f := range fields {
		refs[i] = MetricRef{TimeField: data.NewField("time", nil, []time.Time{{}}), ValueField: f}
	}
	return refs
}

func randomFields(rng *rand.Rand, n int) []*data.Field {
	names := []string{"value", "value", "cpu", "mem"}
	fields := make([]*data.Field, n)
	for i := range fields {
		var labels data.Labels
		switch rng.IntN(10) {
		case 0:
		case 1:
			labels = data.Labels{}
		default:
			labels = data.Labels{}
			for k := range 1 + rng.IntN(4) {
				labels[fmt.Sprintf("k%d", k)] = fmt.Sprintf("v%d", rng.IntN(20))
			}
		}
		fields[i] = data.NewField(names[rng.IntN(len(names))], labels, []float64{float64(i)})
	}
	return fields
}
