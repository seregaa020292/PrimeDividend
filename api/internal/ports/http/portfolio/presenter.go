package portfolio

import "primedivident/internal/modules/portfolio/entity"

func presenterPortfolio(portfolio entity.Portfolio) Portfolio {
	return Portfolio{
		Id:        portfolio.ID,
		CreatedAt: portfolio.CreatedAt,
	}
}
