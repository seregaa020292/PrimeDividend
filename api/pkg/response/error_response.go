package response

import (
	"net/http"

	"primedivident/pkg/errorn"
)

type ErrorRespond struct {
	Data       *any            `json:"data"`
	Errors     []ErrorResponse `json:"errors"`
	httpStatus int
}

type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func NewErrorRespond(httpStatus int, e errorn.Errorn) ErrorRespond {
	return ErrorRespond{
		Errors:     makeErrors(e),
		httpStatus: httpStatus,
	}
}

func (e ErrorRespond) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}

func BadRequest(e errorn.Errorn) ErrorRespond {
	return NewErrorRespond(http.StatusBadRequest, e)
}

func Unauthorised(e errorn.Errorn) ErrorRespond {
	return NewErrorRespond(http.StatusUnauthorized, e)
}

func Forbidden(e errorn.Errorn) ErrorRespond {
	return NewErrorRespond(http.StatusForbidden, e)
}

func NotFound(e errorn.Errorn) ErrorRespond {
	return NewErrorRespond(http.StatusNotFound, e)
}

func UnprocessableEntity(e errorn.Errorn) ErrorRespond {
	return NewErrorRespond(http.StatusUnprocessableEntity, e)
}

func InternalError(e error) ErrorRespond {
	return NewErrorRespond(
		http.StatusInternalServerError,
		errorn.Unknown(errorn.Message{Error: e}).(errorn.Errorn),
	)
}

func FindErrorType(e errorn.Errorn) ErrorRespond {
	switch e.ErrorType() {
	case errorn.ErrorTypeIncorrectInput:
		return BadRequest(e)
	case errorn.ErrorTypeAuthorization:
		return Unauthorised(e)
	case errorn.ErrorTypeForbidden:
		return Forbidden(e)
	case errorn.ErrorTypeNotFound:
		return NotFound(e)
	}

	return InternalError(e)
}

func makeErrors(e errorn.Errorn) []ErrorResponse {
	errorResponse := make([]ErrorResponse, len(e.Messages()))

	for i, message := range e.Messages() {
		errorResponse[i] = ErrorResponse{
			Message: message.Error.Error(),
			Field:   message.Field,
		}
	}

	return errorResponse
}
