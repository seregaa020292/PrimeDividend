package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

// Example: https://github.com/henrmota/errors-handling-example

const (
	NoType ErrorType = iota
	BadRequest
	NotFound
	Unauthorized
	Forbidden
	Conflict
	ServerError
)

type (
	// ErrorType тип ошибки
	ErrorType   uint
	customError struct {
		errorType     ErrorType
		messageError  string
		originalError error
		context       []errorContext
	}
	errorContext struct {
		Field   string
		Message string
	}
)

// New создает новый customError
func (errorType ErrorType) New(msg string) error {
	return errorType.Newf(msg)
}

// Newf создает новый customError с отформатированным сообщением
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{
		errorType:     errorType,
		messageError:  fmt.Sprintf(msg, args...),
		originalError: fmt.Errorf(msg, args...),
		context:       []errorContext{},
	}
}

// Wrap создает новую обернутую ошибку
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrapf создает новую обернутую ошибку с отформатированным сообщением
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{
		errorType:     errorType,
		messageError:  fmt.Sprintf(msg, args...),
		originalError: errors.Wrapf(err, msg, args...),
		context:       []errorContext{},
	}
}

// Error возвращает сообщение customError
func (error customError) Error() string {
	return error.originalError.Error()
}

// Cause возвращает оригинальную ошибку
func (error customError) Cause() error {
	return error.originalError
}

// New создает ошибку без типа
func New(msg string) error {
	return Newf(msg)
}

// Newf создает ошибку без типа с отформатированным сообщением
func Newf(msg string, args ...interface{}) error {
	return customError{
		errorType:     NoType,
		messageError:  fmt.Sprintf(msg, args...),
		originalError: errors.New(fmt.Sprintf(msg, args...)),
		context:       []errorContext{},
	}
}

// Wrap обернуть ошибку строкой
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Wrapf обернуть ошибку строкой формата
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)

	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			messageError:  customErr.messageError,
			originalError: wrappedError,
			context:       customErr.context,
		}
	}

	return customError{
		errorType:     NoType,
		messageError:  fmt.Sprintf(msg, args...),
		originalError: wrappedError,
	}
}

// Cause вернуть исходную ошибку
func Cause(err error) error {
	return errors.Cause(err)
}

// AddErrorContext добавляет контекст к ошибке
func AddErrorContext(err error, field, message string) error {
	context := errorContext{
		Field:   field,
		Message: message,
	}

	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			messageError:  customErr.messageError,
			originalError: customErr.originalError,
			context:       append(customErr.context, context),
		}
	}

	return customError{
		errorType:     NoType,
		messageError:  err.Error(),
		originalError: err,
		context:       append([]errorContext{}, context),
	}
}

// GetErrorContext возвращает контекст ошибки
func GetErrorContext(err error) []errorContext {
	if customErr, ok := err.(customError); ok {
		return customErr.context
	}

	return []errorContext{}
}

// GetMessage возвращает оригинальную ошибку
func GetMessage(err error) string {
	if customErr, ok := err.(customError); ok {
		return customErr.messageError
	}

	return err.Error()
}

// GetType возвращает тип ошибки
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}

// IsCustom проверяет на кастомный тип ошибки
func IsCustom(err error) bool {
	_, ok := err.(customError)

	return ok
}
