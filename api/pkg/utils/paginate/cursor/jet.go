package cursor

import (
	"fmt"

	jet "github.com/go-jet/jet/v2/postgres"
)

type Jet struct {
	input PaginateInput
	id    jet.ColumnString
	time  jet.ColumnTimestampz
}

func NewJet(input PaginateInput, id jet.ColumnString, time jet.ColumnTimestampz) Jet {
	return Jet{
		input: input,
		id:    id,
		time:  time,
	}
}

func (j Jet) PagingSetting(stmt jet.SelectStatement, andCondition jet.BoolExpression) jet.SelectStatement {
	condition := j.GetCondition()

	if andCondition != nil {
		condition = condition.AND(andCondition)
	}

	return stmt.
		WHERE(condition).
		ORDER_BY(j.GetOrderBy()).
		LIMIT(j.GetLimit())
}

func (j Jet) GetCondition() jet.BoolExpression {
	if j.input.Cursor.IsEmpty() {
		return jet.Bool(true)
	}

	return jet.BoolExp(jet.Raw(
		fmt.Sprintf(
			"(%[1]s.%[2]s, %[1]s.%[3]s) %[4]s (#time, #id)",
			j.id.TableName(), j.time.Name(), j.id.Name(), j.input.Cursor.Operator(),
		),
		jet.RawArgs{
			"#time": j.input.Cursor.Time.Local(),
			"#id":   j.input.Cursor.ID.String(),
		},
	))
}

func (j Jet) GetOrderBy() (jet.OrderByClause, jet.OrderByClause) {
	if j.input.Cursor.HasAndPrev() {
		return j.time.ASC(), j.id.ASC()
	}
	return j.time.DESC(), j.id.DESC()
}

func (j Jet) GetLimit() int64 {
	return int64(j.input.GetLimitOver())
}
