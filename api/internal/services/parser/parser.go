package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"primedividend/api/internal/config"
	"primedividend/api/internal/models/app/public/model"
	currencyRepo "primedividend/api/internal/modules/currency/repository"
	instrumentRepo "primedividend/api/internal/modules/instrument/repository"
	marketRepo "primedividend/api/internal/modules/market/repository"
	"primedividend/api/internal/modules/market/service/quotes/providers"
	providerRepo "primedividend/api/internal/modules/provider/repository"
	registerRepo "primedividend/api/internal/modules/register/repository"
	"primedividend/api/pkg/utils/errlog"
)

type Parser struct {
	config         config.Tinkoff
	instrumentRepo instrumentRepo.Repository
	currencyRepo   currencyRepo.Repository
	providerRepo   providerRepo.Repository
	marketRepo     marketRepo.Repository
	registerRepo   registerRepo.Repository
	provider       model.Providers
	instrumentsMap map[string]uuid.UUID
	currenciesMap  map[string]uuid.UUID
}

func NewParser(
	config config.Tinkoff,
	instrumentRepo instrumentRepo.Repository,
	currencyRepo currencyRepo.Repository,
	providerRepo providerRepo.Repository,
	marketRepo marketRepo.Repository,
	registerRepo registerRepo.Repository,
) Parser {
	return Parser{
		config:         config,
		instrumentRepo: instrumentRepo,
		currencyRepo:   currencyRepo,
		providerRepo:   providerRepo,
		marketRepo:     marketRepo,
		registerRepo:   registerRepo,
		instrumentsMap: make(map[string]uuid.UUID),
		currenciesMap:  make(map[string]uuid.UUID),
	}
}

type (
	stock struct {
		Figi              string  `json:"figi"`
		Ticker            string  `json:"ticker"`
		Isin              string  `json:"isin"`
		MinPriceIncrement float32 `json:"minPriceIncrement"`
		Lot               int     `json:"lot"`
		Currency          string  `json:"currency"`
		Name              string  `json:"name"`
		Type              string  `json:"type"`
	}
	responseStock struct {
		TrackingId string `json:"trackingId"`
		Status     string `json:"status"`
		Payload    struct {
			Total       int     `json:"total"`
			Instruments []stock `json:"instruments"`
		} `json:"payload"`
	}
)

func (p *Parser) Select() error {
	instruments, err := p.instrumentRepo.GetAll()
	if err != nil {
		return err
	}

	currencies, err := p.currencyRepo.GetAll()
	if err != nil {
		return err
	}

	p.provider, err = p.providerRepo.GetByTitle(providers.TinkoffProvider)
	if err != nil {
		return err
	}

	for _, v := range instruments {
		p.instrumentsMap[v.Title] = v.ID
	}

	for _, v := range currencies {
		p.currenciesMap[v.Title] = v.ID
	}

	return nil
}

func (p *Parser) Execute(instrument string) error {
	var responseStock responseStock

	log.Printf("start parse %s...", instrument)
	defer log.Printf("end parse %s", instrument)

	if err := p.httpRequest(&responseStock, instrument); err != nil {
		return err
	}

	for _, stock := range responseStock.Payload.Instruments {
		market, err := p.marketRepo.Add(model.Markets{
			Title:        stock.Name,
			Ticker:       stock.Ticker,
			ImageURL:     nil,
			CurrencyID:   p.currenciesMap[stock.Currency],
			InstrumentID: p.instrumentsMap[strings.ToUpper(stock.Type)],
		})
		if err != nil {
			if strings.Contains(err.Error(), "23505") {
				fmt.Println("DUPLICATE:", stock)
				continue
			} else {
				return err
			}
		}

		if _, err := p.registerRepo.Add(model.Registers{
			Identify:   stock.Figi,
			ProviderID: p.provider.ID,
			MarketID:   market.ID,
		}); err != nil {
			if strings.Contains(err.Error(), "23505") {
				fmt.Println("DUPLICATE:", stock)
			} else {
				return err
			}
		}
	}

	return nil
}

func (p *Parser) httpRequest(body any, entity string) error {
	url := fmt.Sprintf("https://api-invest.tinkoff.ru/openapi/sandbox/market/%s", entity)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.config.AuthToken))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer errlog.FnPrintln(response.Body.Close)

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error response: %s", response.Status)
	}

	return json.NewDecoder(response.Body).Decode(body)
}
