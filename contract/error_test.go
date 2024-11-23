package contract_test

import (
	"testing"

	c "github.com/grafana/dataplane/contract"
	"github.com/stretchr/testify/require"
)

func TestContractErrors_Error(t *testing.T) {
	t.Run("nil ContractErrors", func(t *testing.T) {
		var contractErrors *c.ContractErrors
		require.Equal(t, "", contractErrors.Error())
	})

	t.Run("empty ContractErrors", func(t *testing.T) {
		contractErrors := c.ContractErrors{}
		require.Equal(t, "", contractErrors.Error())
	})

	t.Run("single ContractError", func(t *testing.T) {
		contractErrors := c.ContractErrors{c.ContractError{Message: "error 1"}}
		require.Equal(t, "error 1", contractErrors.Error())
	})

	t.Run("multiple ContractErrors", func(t *testing.T) {
		contractErrors := c.ContractErrors{
			{Message: "error 1"},
			{Message: "error 2"},
		}
		require.Equal(t, contractErrors.Error(), "error 1\nerror 2")
	})

	t.Run("multiple ContractErrors with null errors", func(t *testing.T) {
		contractErrors := c.ContractErrors{
			{Message: "error 1"},
			{},
			{Message: "error 2"},
		}
		require.Equal(t, contractErrors.Error(), "error 1\n\nerror 2")
	})
}

func TestContractError_Error(t *testing.T) {
	t.Run("ContractError message", func(t *testing.T) {
		contractError := c.ContractError{Message: "test error"}
		require.Equal(t, "test error", contractError.Error())
	})
}
