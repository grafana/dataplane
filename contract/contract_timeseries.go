package contract

import "github.com/grafana/grafana-plugin-sdk-go/data"

var (
	// TimeseriesWide_0_1 represents a stable contract for TimeSeriesWide frame type version 0.1.
	TimeseriesWide_0_1 Contract = Contract{
		FrameType:     data.FrameTypeTimeSeriesWide,
		Version:       data.FrameTypeVersion{0, 1},
		ContractStage: ContractStageStable,
		Rules:         Rules{},
	}
	// TimeseriesMulti_0_1 represents a stable contract for TimeSeriesMulti frame type version 0.1.
	TimeseriesMulti_0_1 Contract = Contract{
		FrameType:     data.FrameTypeTimeSeriesMulti,
		Version:       data.FrameTypeVersion{0, 1},
		ContractStage: ContractStageStable,
		Rules:         Rules{},
	}
	// TimeseriesLong_0_1 represents a stable contract for TimeSeriesLong frame type version 0.1.
	TimeseriesLong_0_1 Contract = Contract{
		FrameType:     data.FrameTypeTimeSeriesLong,
		Version:       data.FrameTypeVersion{0, 1},
		ContractStage: ContractStageStable,
		Rules:         Rules{},
	}
)
