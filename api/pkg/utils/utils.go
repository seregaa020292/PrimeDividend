package utils

func NewOf[T any](value T) *T {
	return &value
}
