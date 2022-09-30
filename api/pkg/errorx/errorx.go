package errorx

import (
	"fmt"

	"github.com/pkg/errors"
)

// Example: https://github.com/henrmota/errors-handling-example

const (
	NoType ErrorType = iota
	BadRequest
	NotFound
)

type (
	// ErrorType тип ошибки
	ErrorType   uint
	customError struct {
		errorType     ErrorType
		originalError error
		context       errorContext
	}
	errorContext struct {
		Field   string
		Message string
	}
)

// New создает новую customError
func (errorType ErrorType) New(msg string) error {
	return customError{
		errorType:     errorType,
		originalError: errors.New(msg),
	}
}

// Newf создает новый customError с отформатированным сообщением
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{
		errorType:     errorType,
		originalError: fmt.Errorf(msg, args...),
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
		originalError: errors.Wrapf(err, msg, args...),
	}
}

// Error возвращает сообщение customError
func (error customError) Error() string {
	return error.originalError.Error()
}

// New создает ошибку без типа
func New(msg string) error {
	return customError{
		errorType:     NoType,
		originalError: errors.New(msg),
	}
}

// Newf создает ошибку без типа с отформатированным сообщением
func Newf(msg string, args ...interface{}) error {
	return customError{
		errorType:     NoType,
		originalError: errors.New(fmt.Sprintf(msg, args...)),
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
			originalError: wrappedError,
			context:       customErr.context,
		}
	}

	return customError{
		errorType:     NoType,
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
			originalError: customErr.originalError,
			context:       context,
		}
	}

	return customError{
		errorType:     NoType,
		originalError: err,
		context:       context,
	}
}

// GetErrorContext возвращает контекст ошибки
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}

	if customErr, ok := err.(customError); ok || customErr.context != emptyContext {
		return map[string]string{
			"field":   customErr.context.Field,
			"message": customErr.context.Message,
		}
	}

	return nil
}

// GetType возвращает тип ошибки
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}
