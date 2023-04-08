package exchangerate

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const defaultCurrencyCode = "JPY"

var ErrConvert = errors.New("convert error")

type API struct {
	apiKey  string
	baseURL *url.URL
}

func New(baseURL, apiKey string) (*API, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &API{
		apiKey:  apiKey,
		baseURL: u,
	}, nil
}

type (
	convertResponse struct {
		Success bool    `json:"success"`
		Result  float64 `json:"result"`
		Date    string  `json:"date"`
		Query   query   `json:"query"`
		Info    info    `json:"info"`
	}

	query struct {
		To     string `json:"to"`
		From   string `json:"from"`
		Amount int    `json:"amount"`
	}

	info struct {
		Rate      float64 `json:"rate"`
		Timestamp int     `json:"timestamp"`
	}
)

type exchangeRateException struct {
	Message string `json:"message"`
}

func (a API) Convert(currencyCode string) (float64, error) {
	u := a.baseURL.JoinPath("exchangerates_data", "convert")

	v := url.Values{}
	v.Add("to", currencyCode)
	v.Add("from", defaultCurrencyCode)
	v.Add("amount", "1")

	u.RawQuery = v.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("apikey", a.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var e exchangeRateException
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("%w: status code=%d, message=%s", ErrConvert, resp.StatusCode, e.Message)
	}

	var r convertResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}

	return r.Info.Rate, nil
}
