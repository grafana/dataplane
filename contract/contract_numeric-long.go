package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var (
	// Contract_NumericLong_Version_0_1 represents a stable contract for NumericLong frame type version 0.1.
	Contract_NumericLong_Version_0_1 Contract = New(
		data.FrameTypeNumericLong,
		data.FrameTypeVersion{0, 1},
		ContractStageStable, rules.Rules{
			rules.CheckAtLeastOneFrame,
			//TODO: more rules to be added
		},
	)
)
