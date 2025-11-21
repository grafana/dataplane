package contract

import (
	"github.com/grafana/dataplane/contract/rules"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var (
	// Contract_DirectoryListing_Version_Unknown represents DirectoryListing frame type contract with unknown version.
	Contract_DirectoryListing_Version_Unknown Contract = New(
		data.FrameTypeDirectoryListing,
		data.FrameTypeVersion{0, 0},
		ContractStageUnknown,
		rules.Rules{
			rules.CheckAtLeastOneFrame,
			//TODO: more rules to be added
		},
	)
)
