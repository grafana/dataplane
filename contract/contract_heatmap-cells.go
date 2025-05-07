package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var (
	// Contract_HeatmapCells_Version_Unstable represents an unstable contract for heatmap-cells frame type.
	Contract_HeatmapCells_Version_Unstable Contract = New(
		data.FrameType("heatmap-cells"),
		data.FrameTypeVersion{0, 0},
		ContractStageUnstable,
		rules.Rules{
			rules.CheckAtLeastOneFrame,
			//TODO: more rules to be added
		},
	)
)
