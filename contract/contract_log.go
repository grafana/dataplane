package contract

import "github.com/grafana/grafana-plugin-sdk-go/data"

var (
	// LogLinesContract_unstable represents an unstable contract for LogLines frame type.
	LogLinesContract_unstable Contract = Contract{
		FrameType:     data.FrameTypeLogLines,
		Version:       data.FrameTypeVersion{0, 0},
		ContractStage: ContractStageUnstable,
		Rules:         Rules{},
	}
)
