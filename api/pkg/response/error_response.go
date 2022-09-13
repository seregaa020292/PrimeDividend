package response

import (
	"net/http"

	"primedivident/pkg/errorn"
)

type ErrorResponse struct {
	Data  *any  `json:"data"`
	Error error `json:"error"`

	status int
}

func NewByError(err error) ErrorResponse {
	e, ok := err.(errorn.Error)
	if !ok {
		e = errorn.ErrorUnknown.Wrap(e)
	}

	return ErrorResponse{
		Error:  e,
		status: e.Status(),
	}
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.status)
	return nil
}
