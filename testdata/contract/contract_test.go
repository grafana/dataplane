package contract_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/grafana/dataplane/testdata/contract"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/stretchr/testify/require"
)

func TestGetExamples(t *testing.T) {
	_, err := contract.GetExamples()
	require.NoError(t, err)
}

func TestValidExamples(t *testing.T) {
	examples, err := contract.GetExamples()
	require.NoError(t, err)
	require.Greater(t, len(examples), 1)

	t.Run("all sumaries must end in period", func(t *testing.T) {
		for _, e := range examples {
			info := e.GetInfo()
			require.True(t, strings.HasSuffix(info.Summary, "."), fmt.Sprintf("Summary: %q Path: %q", info.Summary, info.Path))
		}
	})
	t.Run("all sumaries must have collectionVersion >= 1", func(t *testing.T) {
		for _, e := range examples {
			info := e.GetInfo()
			require.GreaterOrEqual(t, e.GetInfo().CollectionVersion, 1, info.Path)
		}
	})
}

func TestExamplesFilter(t *testing.T) {
	e, err := contract.GetExamples()
	require.NoError(t, err)

	t.Run("frametype and kind must match if not zero value", func(t *testing.T) {
		_, err := e.Filter(contract.FilterOptions{Type: data.FrameTypeTimeSeriesLong})
		require.Error(t, err)
	})

	t.Run("version filter will error for no results for version with no match", func(t *testing.T) {
		_, err := e.Filter(contract.FilterOptions{Version: data.FrameTypeVersion{124, 1233}})
		require.Error(t, err)
	})

	t.Run("can filter to numeric long", func(t *testing.T) {
		numLongExamples, err := e.Filter(contract.FilterOptions{Type: data.FrameTypeNumericLong})

		require.NoError(t, err)
		for _, e := range numLongExamples {
			require.Equal(t, data.FrameTypeNumericLong, e.GetInfo().Type)
		}
	})
}
