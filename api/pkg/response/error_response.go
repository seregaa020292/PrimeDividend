package response

import (
	"net/http"
	"primedivident/pkg/errorn"
)

type ErrorRespond struct {
	Data       *interface{}    `json:"data"`
	Errors     []ErrorResponse `json:"errors"`
	httpStatus int
}

type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

func (e ErrorRespond) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}

func BadRequest(e errorn.Errorn) ErrorRespond {
	return ErrorRespond{
		Errors:     makeErrors(e),
		httpStatus: http.StatusBadRequest,
	}
}

func Unauthorised(e errorn.Errorn) ErrorRespond {
	return ErrorRespond{
		Errors:     makeErrors(e),
		httpStatus: http.StatusUnauthorized,
	}
}

func Forbidden(e errorn.Errorn) ErrorRespond {
	return ErrorRespond{
		Errors:     makeErrors(e),
		httpStatus: http.StatusForbidden,
	}
}

func NotFound(e errorn.Errorn) ErrorRespond {
	return ErrorRespond{
		Errors:     makeErrors(e),
		httpStatus: http.StatusNotFound,
	}
}

func UnprocessableEntity(e errorn.Errorn) ErrorRespond {
	return ErrorRespond{
		Errors:     makeErrors(e),
		httpStatus: http.StatusUnprocessableEntity,
	}
}

func InternalError(e error) ErrorRespond {
	err := errorn.Unknown(errorn.Message{Error: e}).(errorn.Errorn)
	return ErrorRespond{
		Errors:     makeErrors(err),
		httpStatus: http.StatusInternalServerError,
	}
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
