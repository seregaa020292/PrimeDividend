package response

import (
	"net/http"
	"primedivident/internal/mistake"
)

type ErrorResponse struct {
	Code    string
	Message string
}

type ErrorRespond struct {
	Data       *interface{}  `json:"data"`
	Err        ErrorResponse `json:"error"`
	httpStatus int
}

func (e ErrorRespond) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}

func InternalError(m string) ErrorRespond {
	return ErrorRespond{
		httpStatus: http.StatusInternalServerError,
		Err: ErrorResponse{
			Code:    "",
			Message: m,
		},
	}
}

func BadRequest(m string) ErrorRespond {
	return ErrorRespond{
		httpStatus: http.StatusBadRequest,
		Err: ErrorResponse{
			Code:    "",
			Message: m,
		},
	}
}

func NotFound(m string) ErrorRespond {
	return ErrorRespond{
		httpStatus: http.StatusNotFound,
		Err: ErrorResponse{
			Code:    "",
			Message: m,
		},
	}
}

func Unauthorised(m string) ErrorRespond {
	return ErrorRespond{
		httpStatus: http.StatusUnauthorized,
		Err: ErrorResponse{
			Code:    "",
			Message: m,
		},
	}
}

func FindErrorType(slugError mistake.SlugError) ErrorRespond {
	switch slugError.ErrorType() {
	case mistake.ErrorTypeAuthorization:
		return Unauthorised(slugError.Slug())
	case mistake.ErrorTypeIncorrectInput:
		return BadRequest(slugError.Slug())
	}

	return InternalError(slugError.Slug())
}
