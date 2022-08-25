package response

import (
	"net/http"
	"primedivident/internal/errors"
)

type Error struct {
	Code    string
	Message string
}

type ErrorResponse struct {
	Data       *interface{} `json:"data"`
	Err        Error        `json:"error"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}

func InternalError(m string) ErrorResponse {
	return ErrorResponse{
		httpStatus: http.StatusInternalServerError,
		Err: Error{
			Code:    "",
			Message: m,
		},
	}
}

func BadRequest(m string) ErrorResponse {
	return ErrorResponse{
		httpStatus: http.StatusBadRequest,
		Err: Error{
			Code:    "",
			Message: m,
		},
	}
}

func NotFound(m string) ErrorResponse {
	return ErrorResponse{
		httpStatus: http.StatusNotFound,
		Err: Error{
			Code:    "",
			Message: m,
		},
	}
}

func Unauthorised(m string) ErrorResponse {
	return ErrorResponse{
		httpStatus: http.StatusUnauthorized,
		Err: Error{
			Code:    "",
			Message: m,
		},
	}
}

func FindErrorType(slugError errors.SlugError) ErrorResponse {
	switch slugError.ErrorType() {
	case errors.ErrorTypeAuthorization:
		return Unauthorised(slugError.Slug())
	case errors.ErrorTypeIncorrectInput:
		return BadRequest(slugError.Slug())
	}

	return InternalError(slugError.Slug())
}
