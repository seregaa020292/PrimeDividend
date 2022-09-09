package portfolio

import (
	"primedivident/internal/infrastructure/http/openapi"
	"primedivident/internal/modules/portfolio/entity"
)

func presenterPortfolio(portfolio entity.Portfolio) openapi.Portfolio {
	return openapi.Portfolio{
		Id:        portfolio.ID,
		CreatedAt: portfolio.CreatedAt,
	}
}
