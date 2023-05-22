package transformations

import (
	"errors"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type LimitOptions struct {
	LimitField int `json:"limitField,omitempty"`
}

func Limit(input []*data.Frame, options LimitOptions) ([]*data.Frame, error) {
	return nil, errors.New("limit transformation not implemented yet")
}
