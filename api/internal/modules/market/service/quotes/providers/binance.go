package providers

import "github.com/adshao/go-binance/v2"

type Binance struct {
	*binance.Client
}

func NewBinance(apiKey, secretKey string) Binance {
	return Binance{
		Client: binance.NewClient(apiKey, secretKey),
	}
}
