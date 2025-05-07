package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var (
	// Contract_TimeseriesLong_Version_0_1 represents a stable contract for TimeSeriesLong frame type version 0.1.
	Contract_TimeseriesLong_Version_0_1 Contract = New(
		data.FrameTypeTimeSeriesLong,
		data.FrameTypeVersion{0, 1},
		ContractStageStable, rules.Rules{
			rules.CheckAtLeastOneFrame,
			//TODO: more rules to be added
		},
	)
)
