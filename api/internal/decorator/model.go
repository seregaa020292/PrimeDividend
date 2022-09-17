package decorator

type Model[Q any] struct {
	Table   string
	Columns Q
}
