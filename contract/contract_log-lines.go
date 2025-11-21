package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var (
	// Contract_LogLines_Version_Unstable represents an unstable contract for LogLines frame type.
	Contract_LogLines_Version_Unstable Contract = New(
		data.FrameTypeLogLines,
		data.FrameTypeVersion{0, 0},
		ContractStageUnstable,
		rules.Rules{
			rules.CheckAtLeastOneFrame,
			//TODO: more rules to be added
		},
	)
)
