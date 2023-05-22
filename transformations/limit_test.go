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
		// TODO: Add test cases.
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
