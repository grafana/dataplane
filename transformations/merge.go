package transformations

import (
	"errors"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type MergeOptions struct {
}

func Merge(input []*data.Frame, options MergeOptions) (*data.Frame, error) {
	return nil, errors.New("merge transformation not implemented yet")
}
