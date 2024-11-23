package contract

import "github.com/grafana/grafana-plugin-sdk-go/data"

var (
	// NumericWide_0_1 represents a stable contract for NumericWide frame type version 0.1.
	NumericWide_0_1 Contract = Contract{
		FrameType:     data.FrameTypeNumericWide,
		Version:       data.FrameTypeVersion{0, 1},
		ContractStage: ContractStageStable,
		Rules:         Rules{},
	}
	// NumericMulti_0_1 represents a stable contract for NumericMulti frame type version 0.1.
	NumericMulti_0_1 Contract = Contract{
		FrameType:     data.FrameTypeNumericMulti,
		Version:       data.FrameTypeVersion{0, 1},
		ContractStage: ContractStageStable,
		Rules:         Rules{},
	}
	// NumericLong_0_1 represents a stable contract for NumericLong frame type version 0.1.
	NumericLong_0_1 Contract = Contract{
		FrameType:     data.FrameTypeNumericLong,
		Version:       data.FrameTypeVersion{0, 1},
		ContractStage: ContractStageStable,
		Rules:         Rules{},
	}
)
