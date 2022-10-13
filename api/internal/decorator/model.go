package decorator

type ColumnAssigment[T any] []T

func (c ColumnAssigment[T]) Column() T {
	return c[0]
}

func (c ColumnAssigment[T]) ColumnList() []T {
	if len(c) > 0 {
		return c[1:]
	}
	return []T{}
}
