package graphql

import (
	"github.com/ahmadmuzakki/graphql/gqlerrors"
)

// type Schema interface{}

type Result struct {
	Data   interface{}                `json:"data"`
	Errors []gqlerrors.FormattedError `json:"errors,omitempty"`
}

func (r *Result) HasErrors() bool {
	return len(r.Errors) > 0
}
