package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var (
	// Contract_Table_Version_Unknown represents an unknown contract for Table frame type.
	Contract_Table_Version_Unknown Contract = New(
		data.FrameTypeTable,
		data.FrameTypeVersion{0, 0},
		ContractStageUnknown,
		rules.Rules{
			rules.CheckAtLeastOneFrame,
			//TODO: more rules to be added
		},
	)
)
