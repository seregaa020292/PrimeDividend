package cursor

import (
	"primedividend/api/internal/config/consts"
	"primedividend/api/pkg/utils/gog"
)

type PaginateInput struct {
	Limit  int
	Cursor Cursor
}

func NewPaginateInput(limit *int, cursor *string) (PaginateInput, error) {
	cursorDirection, err := cursorInput(cursor)
	if err != nil {
		return PaginateInput{}, err
	}

	return PaginateInput{
		Limit:  limitInput(limit),
		Cursor: cursorDirection,
	}, nil
}

func (p PaginateInput) GetLimitOver() int {
	return gog.If(p.Limit > 0, p.Limit+1, 0)
}

func (p PaginateInput) EqLimitOver(length int) bool {
	return p.GetLimitOver() == length
}

func limitInput(limit *int) int {
	if limit == nil {
		return consts.PageLimitDefault
	}

	limitNum := *limit

	if limitNum <= 0 {
		return 0
	}

	if limitNum > consts.PageLimitMax {
		return consts.PageLimitMax
	}

	return limitNum
}

func cursorInput(cursor *string) (Cursor, error) {
	if cursor == nil {
		return Cursor{}, nil
	}

	return Decode(*cursor)
}
