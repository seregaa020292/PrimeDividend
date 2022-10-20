package offset

import (
	"math"

	"primedivident/internal/config/consts"
)

type Pager struct {
	CurrentPage int  `json:"currentPage"`
	MaxPage     int  `json:"maxPage"`
	Total       int  `json:"total"`
	Start       *int `json:"-"`
	End         *int `json:"-"`
}

func Paging(recordCount, size, page int) Pager {
	maxPage := consts.PageLimitDefault

	if recordCount > 0 {
		maxPage = int(math.Ceil(float64(recordCount) / float64(size)))
	}

	p := Pager{
		CurrentPage: page,
		MaxPage:     maxPage,
		Total:       recordCount,
	}

	start := size * (page - 1)
	end := start + size

	if start < recordCount {
		p.Start = &start
	}

	if end < recordCount {
		p.End = &end
	}

	return p
}
