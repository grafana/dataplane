package contract

import "github.com/grafana/grafana-plugin-sdk-go/data"

var (
	// TableContract_unknown represents an unknown contract for Table frame type.
	TableContract_unknown Contract = Contract{
		FrameType:     data.FrameTypeTable,
		Version:       data.FrameTypeVersion{0, 0},
		ContractStage: ContractStageUnknown,
		Rules:         Rules{},
	}
)
