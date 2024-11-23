package contract

import "github.com/grafana/grafana-plugin-sdk-go/data"

var (
	// HeatmapRowsContract_unstable represents an unstable contract for heatmap-rows frame type.
	HeatmapRowsContract_unstable Contract = Contract{
		FrameType:     data.FrameType("heatmap-rows"),
		Version:       data.FrameTypeVersion{0, 0},
		ContractStage: ContractStageUnstable,
		Rules:         Rules{},
	}
	// HeatmapCellsContract_unstable represents an unstable contract for heatmap-cells frame type.
	HeatmapCellsContract_unstable Contract = Contract{
		FrameType:     data.FrameType("heatmap-cells"),
		Version:       data.FrameTypeVersion{0, 0},
		ContractStage: ContractStageUnstable,
		Rules:         Rules{},
	}
)
