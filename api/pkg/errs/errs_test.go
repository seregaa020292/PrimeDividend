package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	err := BadRequest.New("an_error")
	errWithContext := AddErrorContext(err, "a_field", "the field is empty")

	expectedContext := []errorContext{{Field: "a_field", Message: "the field is empty"}}

	assert.Equal(t, NoType, GetType(errWithContext))
	assert.ElementsMatch(t, expectedContext, GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestContextInNoTypeError(t *testing.T) {
	err := New("a custom error")

	errWithContext := AddErrorContext(err, "a_field", "the field is empty")

	expectedContext := []errorContext{{Field: "a_field", Message: "the field is empty"}}

	assert.Equal(t, NoType, GetType(errWithContext))
	assert.ElementsMatch(t, expectedContext, GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestWrapf(t *testing.T) {
	err := New("an_error")
	wrappedError := BadRequest.Wrapf(err, "error %s", "1")

	assert.Equal(t, BadRequest, GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error")
}

func TestWrapfInNoTypeError(t *testing.T) {
	err := Newf("an_error %s", "2")
	wrappedError := Wrapf(err, "error %s", "1")

	assert.Equal(t, NoType, GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error 2")
}
