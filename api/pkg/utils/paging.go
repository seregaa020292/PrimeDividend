package utils

import "math"

type Pager struct {
	CurrentPage int  `json:"currentPage"`
	MaxPage     int  `json:"maxPage"`
	Start       *int `json:"-"`
	End         *int `json:"-"`
}

const DefaultPagerPage = 25

func Paging(listNum, size, page int) Pager {
	maxPage := DefaultPagerPage

	if listNum > 0 {
		maxPage = int(math.Ceil(float64(listNum) / float64(size)))
	}

	p := Pager{
		CurrentPage: page,
		MaxPage:     maxPage,
	}

	start := size * (page - 1)
	end := start + size

	if start < listNum {
		p.Start = &start
	}

	if end < listNum {
		p.End = &end
	}

	return p
}
