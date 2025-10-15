package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var (
	// Contract_NumericMulti_Version_0_1 represents a stable contract for NumericMulti frame type version 0.1.
	Contract_NumericMulti_Version_0_1 Contract = New(
		data.FrameTypeNumericMulti,
		data.FrameTypeVersion{0, 1},
		ContractStageStable, rules.Rules{
			rules.CheckAtLeastOneFrame,
			//TODO: more rules to be added
		},
	)
)
