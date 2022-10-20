package cursor

const (
	NextDirection Direction = true
	PrevDirection Direction = false
)

type Direction bool

func (d Direction) IsNext() bool {
	return d == NextDirection
}

func (d Direction) IsPrev() bool {
	return d == PrevDirection
}
