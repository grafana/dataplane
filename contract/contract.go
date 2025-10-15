package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// Contract represents a data frame contract with a specific frame type, version, stage and corresponding rules.
type Contract struct {
	frameType data.FrameType
	version   data.FrameTypeVersion
	stage     ContractStage
	rules     rules.Rules
}

func New(frameType data.FrameType, version data.FrameTypeVersion, contractStage ContractStage, rules rules.Rules) Contract {
	c := Contract{
		frameType: frameType,
		version:   version,
		stage:     contractStage,
		rules:     rules,
	}
	return c
}

// Compliance checks if the given data frames comply with the contract.
// It returns a boolean indicating compliance and a list of contract errors if any.
func (contract *Contract) Compliance(frames data.Frames) (bool, ContractErrors) {
	errs := ContractErrors{}
	complianceErrors := contract.rules.Compliance(frames)
	for _, v := range complianceErrors {
		if v.Message != "" {
			errs = append(errs, ContractError{Message: v.Message})
		}
	}
	return len(errs) < 1, errs
}
