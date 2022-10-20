package response

import (
	"net/http"

	"primedivident/internal/infrastructure/server/openapi"
	"primedivident/pkg/errs"
	"primedivident/pkg/errs/errmsg"
)

type ErrorResponse struct {
	openapi.Error
	status int
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: openapi.Error{
			Error: openapi.ErrorMessage{
				Details: newDetails(err),
				Message: newMessage(err),
			},
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

	return errmsg.ServerError
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

func newDetails(err error) []openapi.ErrorDetail {
	ctxErrors := errs.GetErrorContext(err)
	details := make([]openapi.ErrorDetail, len(ctxErrors))

	for i, v := range ctxErrors {
		details[i] = openapi.ErrorDetail(v)
	}

	return details
}
