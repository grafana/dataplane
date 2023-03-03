package contract_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/grafana/dataplane/testdata/contract"
	"github.com/stretchr/testify/require"
)

func TestStuff(t *testing.T) {
	t.Run("Can get examples", func(t *testing.T) {
		var err error
		_, err = contract.GetExamples()
		require.NoError(t, err)
	})

	t.Run("All sumarries must end in period", func(t *testing.T) {
		e, err := contract.GetExamples()
		require.NoError(t, err)
		examples := e.GetAllAsList()

		for _, e := range examples {
			info := e.GetInfo()
			require.True(t, strings.HasSuffix(info.Summary, "."), fmt.Sprintf("Summary: %q Path: %q", info.Summary, info.Path))
		}
	})
}
