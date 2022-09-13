package decorator

type Model[Q any] struct {
	Table  string
	Fields Q
}