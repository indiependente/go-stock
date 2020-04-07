package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type GlobalQuoteResponse struct {
	GlobalQuote GlobalQuote `json:"Global Quote"`
}

type Caller interface {
	GlobalQuote(ctx context.Context, symbol string) (GlobalQuoteResponse, error)
}

type Client struct {
	Doer   Doer
	URL    string
	ApiKey string
}

func (c Client) GlobalQuote(ctx context.Context, symbol string) (GlobalQuoteResponse, error) {
	url := fmt.Sprintf("%s/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", c.URL, symbol, c.ApiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("could not create new request: %w", err)
	}

	resp, err := c.Doer.Do(req)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("error in request: %w", resp)
	}
	defer resp.Body.Close() // nolint: errcheck

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("could not read response body: %w", err)
	}

	var gqr GlobalQuoteResponse
	err = json.Unmarshal(data, &gqr)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("could not unmarshal response: %w", err)
	}

	return gqr, nil
}
