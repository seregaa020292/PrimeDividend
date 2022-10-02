package response

import (
	"net/http"

	"primedivident/pkg/errs"
	"primedivident/pkg/errs/bugreport"
)

type (
	errDetail struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
	errResponse struct {
		Message string      `json:"message"`
		Details []errDetail `json:"details"`
	}
	ErrorResponse struct {
		Data  *any        `json:"data"`
		Error errResponse `json:"error"`

		status int
	}
)

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: errResponse{
			Message: newMessage(err),
			Details: newDetails(err),
		},
		status: newStatus(err),
	}
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.status)
	return nil
}

func newMessage(err error) string {
	if errs.IsCustom(err) {
		return errs.GetMessage(err)
	}

	return bugreport.ServerError
}

func newStatus(err error) int {
	switch errs.GetType(err) {
	case errs.BadRequest, errs.NoType:
		return http.StatusBadRequest
	case errs.NotFound:
		return http.StatusNotFound
	case errs.Unauthorized:
		return http.StatusUnauthorized
	case errs.Forbidden:
		return http.StatusForbidden
	case errs.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func newDetails(err error) []errDetail {
	ctxErrors := errs.GetErrorContext(err)
	details := make([]errDetail, len(ctxErrors))

	for i, v := range ctxErrors {
		details[i] = errDetail(v)
	}

	return details
}
