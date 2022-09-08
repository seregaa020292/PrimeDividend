package validator

import (
	"sync"
)

type Validator interface {
	Struct(any) error
	Var(any, string) error
}

var (
	instance Validator
	once     sync.Once
)

func GetValidator() Validator {
	once.Do(func() {
		instance = NewGoPlayground()
	})

	return instance
}
