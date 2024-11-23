package contract

import (
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// Contract represents a data frame contract with a specific frame type, version, stage and corresponding rules.
type Contract struct {
	FrameType     data.FrameType
	Version       data.FrameTypeVersion
	ContractStage ContractStage
	Rules         Rules
}

// Compliance checks if the given data frames comply with the contract.
// It returns a boolean indicating compliance and a list of contract errors if any.
// WARNING: If returned errors contains ErrContractCheckNotImplemented, then skip the validation
func (contract *Contract) Compliance(frames data.Frames) (bool, ContractErrors) {
	errors := ContractErrors{}
	if len(contract.Rules) == 0 {
		errors = append(errors, ErrContractCheckNotImplemented)
		return false, errors
	}
	errors = contract.Rules.IsCompliant(frames)
	return len(errors) < 1, errors
}
