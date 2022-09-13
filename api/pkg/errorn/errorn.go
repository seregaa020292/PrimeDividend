package errorn

import (
	"encoding/json"
	"fmt"

	"primedivident/pkg/utils/gog"
)

type (
	// Error
	//
	//	error - оригинальная ошибка
	//	status - код ошибки HTTP, указанный в ответе
	//	target — идентификатор, который классифицирует ошибку
	//	message - краткое, удобочитаемое сообщение об ошибке
	Error struct {
		error   error
		status  int
		target  Target
		message string
		details []DetailError
	}
	// DetailError
	//
	//	Target – ошибка поля
	//	Message – краткое сообщение, понятное человеку
	DetailError struct {
		Target  string `json:"target"`
		Message string `json:"message"`
	}
	// Target классификация ошибки
	Target int
)

func (e Error) Wrap(err error) Error {
	newErr := e.Clone()
	newErr.error = err
	return newErr
}

func (e Error) Additional(errors ...DetailError) Error {
	newErr := e.Clone()
	newErr.details = append(e.details, errors...)
	return newErr
}

func (e Error) Clone() Error {
	err := NewError(e.target, e.status, e.message, e.details...)
	err.error = e.error
	return err
}

func (e Error) Error() string {
	err := gog.If(e.error != nil, e.error, fmt.Errorf("%s", "_"))

	return fmt.Sprintf(
		"Error: %s; Status: %d; Target: %d; Message: %s; Details: %+v",
		err, e.status, e.target, e.message, e.details,
	)
}

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Status  int           `json:"status"`
		Target  Target        `json:"target"`
		Message string        `json:"message"`
		Details []DetailError `json:"details"`
	}{
		Status:  e.status,
		Target:  e.target,
		Message: e.message,
		Details: e.details,
	})
}

func (e Error) Status() int {
	return e.status
}

func NewError(
	target Target,
	status int,
	message string,
	details ...DetailError,
) Error {
	return Error{
		status:  status,
		target:  target,
		message: message,
		details: append([]DetailError{}, details...),
	}
}
