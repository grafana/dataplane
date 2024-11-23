package contract

import "github.com/grafana/grafana-plugin-sdk-go/data"

var (
	// DirectoryListingContract_unknown represents an unknown contract for DirectoryListing frame type.
	DirectoryListingContract_unknown Contract = Contract{
		FrameType:     data.FrameTypeDirectoryListing,
		Version:       data.FrameTypeVersion{0, 0},
		ContractStage: ContractStageUnknown,
		Rules:         Rules{},
	}
)
