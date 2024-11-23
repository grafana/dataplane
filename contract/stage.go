package contract

// ContractStage defines the stage of a contract.
type ContractStage string

const (
	// ContractStageStable indicates a stable contract.
	ContractStageStable ContractStage = "stable"
	// ContractStageAlpha indicates an alpha stage contract.
	ContractStageAlpha ContractStage = "alpha"
	// ContractStageBeta indicates a beta stage contract.
	ContractStageBeta ContractStage = "beta"
	// ContractStageUnstable indicates an unstable contract.
	// Frame type contract in proposal stage in dataplane monorepo but may/may not available in SDK
	ContractStageUnstable ContractStage = "unstable"
	// ContractStageUnknown indicates an unknown contract.
	// Frame type contract that are not in dataplane monorepo but in SDK
	ContractStageUnknown ContractStage = "unknown"
)
