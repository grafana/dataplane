package transformations_test

import (
	"testing"

	"github.com/grafana/dataplane/transformations"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLimit(t *testing.T) {
	tests := []struct {
		name    string
		input   []*data.Frame
		options transformations.LimitOptions
		want    []*data.Frame
		wantErr error
	}{
		{
			name:  "nil frames array should return back",
			input: []*data.Frame{nil, nil, nil},
			want:  []*data.Frame{nil, nil, nil},
		},
		{
			name:  "empty frame array should return back",
			input: []*data.Frame{},
			want:  []*data.Frame{},
		},
		{
			name: "single frame array should return back single frame",
			input: []*data.Frame{
				data.NewFrame(""),
			},
			want: []*data.Frame{
				data.NewFrame(""),
			},
		},
		{
			name: "multi frame array should return back multi frame with filtered values",
			input: []*data.Frame{
				data.NewFrame("A", data.NewField("f1", nil, []int64{1, 2, 3})),
				data.NewFrame("B", data.NewField("f1", nil, []int64{1, 2})),
				data.NewFrame("C", data.NewField("f1", nil, []int64{})),
				data.NewFrame("D", data.NewField("f1", nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
				data.NewFrame("E", data.NewField("f1", nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})),
			},
			want: []*data.Frame{
				data.NewFrame("A", data.NewField("f1", nil, []int64{1, 2, 3})),
				data.NewFrame("B", data.NewField("f1", nil, []int64{1, 2})),
				data.NewFrame("C", data.NewField("f1", nil, []int64{})),
				data.NewFrame("D", data.NewField("f1", nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
				data.NewFrame("E", data.NewField("f1", nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
			},
		},
		{
			name:    "multi frame array should return back multi frame with filtered values with custom limit speficied",
			options: transformations.LimitOptions{LimitField: 2},
			input: []*data.Frame{
				data.NewFrame("A", data.NewField("f1", nil, []int64{1, 2, 3})),
				data.NewFrame("B", data.NewField("f1", nil, []int64{1, 2})),
				data.NewFrame("C", data.NewField("f1", nil, []int64{})),
				nil,
				data.NewFrame("D", data.NewField("f1", nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
				data.NewFrame("E", data.NewField("f1", nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})),
			},
			want: []*data.Frame{
				data.NewFrame("A", data.NewField("f1", nil, []int64{1, 2})),
				data.NewFrame("B", data.NewField("f1", nil, []int64{1, 2})),
				data.NewFrame("C", data.NewField("f1", nil, []int64{})),
				nil,
				data.NewFrame("D", data.NewField("f1", nil, []int64{1, 2})),
				data.NewFrame("E", data.NewField("f1", nil, []int64{1, 2})),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := transformations.Limit(tt.input, tt.options)
			if tt.wantErr != nil {
				require.NotNil(t, err)
				assert.Equal(t, tt.wantErr, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
