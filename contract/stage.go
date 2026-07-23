package contract

// ContractStage defines the stage of a contract.
type ContractStage string

const (
	// ContractStageStable indicates a stable contract.
	ContractStageStable ContractStage = "stable"
	// ContractStageUnstable indicates an unstable contract.
	// Frame type contract in proposal/alpha/beta stage in dataplane monorepo but may not be defined elsewhere in grafana
	ContractStageUnstable ContractStage = "unstable"
	// ContractStageUnknown indicates an unknown contract.
	// Frame type contract that are not in dataplane monorepo but in SDK or defined elsewhere in grafana orgs repo
	// Some reference are given below
	// https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/types/dataFrameTypes.ts
	// https://github.com/grafana/grafana-plugin-sdk-go/blob/main/data/frame_type.go
	// https://github.com/grafana/plugin-extension-types/blob/main/types/grafana-pluginsplatform-app/index.d.ts
	ContractStageUnknown ContractStage = "unknown"
)
