package exchangerate

import (
	"fmt"
	"net/url"
)

type ExchangeRate struct {
	apiKey  string
	baseURL *url.URL
}

func New(baseURL, apiKey string) (*ExchangeRate, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &ExchangeRate{
		apiKey:  apiKey,
		baseURL: u,
	}, nil
}

func (er ExchangeRate) Convert(currencyCode string) {
	u := er.baseURL.JoinPath("exchangerates_data", "convert")

	fmt.Println(u.String())
}
