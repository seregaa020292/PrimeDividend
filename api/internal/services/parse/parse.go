package parse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"primedivident/internal/config"
	"primedivident/internal/models/app/public/model"
	currencyRepo "primedivident/internal/modules/currency/repository"
	instrumentRepo "primedivident/internal/modules/instrument/repository"
	marketRepo "primedivident/internal/modules/market/repository"
	providerRepo "primedivident/internal/modules/provider/repository"
	registerRepo "primedivident/internal/modules/register/repository"
)

type Parse struct {
	config         config.Tinkoff
	instrumentRepo instrumentRepo.Repository
	currencyRepo   currencyRepo.Repository
	providerRepo   providerRepo.Repository
	marketRepo     marketRepo.Repository
	registerRepo   registerRepo.Repository
}

func NewParse(
	config config.Tinkoff,
	instrumentRepo instrumentRepo.Repository,
	currencyRepo currencyRepo.Repository,
	providerRepo providerRepo.Repository,
	marketRepo marketRepo.Repository,
	registerRepo registerRepo.Repository,
) Parse {
	return Parse{
		config:         config,
		instrumentRepo: instrumentRepo,
		currencyRepo:   currencyRepo,
		providerRepo:   providerRepo,
		marketRepo:     marketRepo,
		registerRepo:   registerRepo,
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

func (p Parse) Execute(instrument string) error {
	type mapUUID map[string]uuid.UUID
	var responseStock responseStock

	if err := p.httpRequest(&responseStock, instrument); err != nil {
		return err
	}

	instruments, err := p.instrumentRepo.GetAll()
	if err != nil {
		return err
	}

	currencies, err := p.currencyRepo.GetAll()
	if err != nil {
		return err
	}

	provider, err := p.providerRepo.GetByTitle("Tinkoff")
	if err != nil {
		return err
	}

	instrumentsMap := make(mapUUID)
	for _, v := range instruments {
		instrumentsMap[v.Title] = v.ID
	}

	currenciesMap := make(mapUUID)
	for _, v := range currencies {
		currenciesMap[v.Title] = v.ID
	}

	for _, stock := range responseStock.Payload.Instruments {
		market, err := p.marketRepo.Add(model.Markets{
			Title:        stock.Name,
			Ticker:       stock.Ticker,
			ImageURL:     nil,
			CurrencyID:   currenciesMap[stock.Currency],
			InstrumentID: instrumentsMap[strings.ToUpper(stock.Type)],
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
			ProviderID: provider.ID,
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

func (p Parse) httpRequest(body any, entity string) error {
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

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error response api-invest.tinkoff: %s", response.Status)
	}

	return json.NewDecoder(response.Body).Decode(body)
}
