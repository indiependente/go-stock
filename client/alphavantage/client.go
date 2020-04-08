package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	urlTemplate = `%s/query?function=%s&symbol=%s&apikey=%s`
)

// Doer defines the behaviour of a component able to perform an HTTP request.
// Returns a response and an error if any.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// GlobalQuoteResponse is the response returned by a GlobalQuote request.
type GlobalQuoteResponse struct {
	GlobalQuote GlobalQuote `json:"Global Quote"`
}

// Caller defines the behaviour of a component able to run requests to Alpha Vantage.
type Caller interface {
	GlobalQuote(ctx context.Context, symbol string) (GlobalQuoteResponse, error)
}

// Client is an Alpha Vintage client.
type Client struct {
	// Doer is an implementation of the Doer interface
	Doer Doer
	// URL is the URL to send requests to.
	URL string
	// APIKey is the key that API requires to identify the user.
	APIKey string
}

// GlobalQuote returns quotes information about the input symbol.
// Returns an error if any.
func (c Client) GlobalQuote(ctx context.Context, symbol string) (GlobalQuoteResponse, error) {
	url := fmt.Sprintf(urlTemplate, "GLOBAL_QUOTE", c.URL, symbol, c.APIKey)

	data, err := c.get(ctx, url)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("could not get global quote: %w", err)
	}

	var gqr GlobalQuoteResponse
	err = json.Unmarshal(data, &gqr)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("could not unmarshal response: %w", err)
	}

	return gqr, nil
}

// WeeklyTimeSeries returns weekly time series for the input symbol for the last 20 years.
// Returns an error if any.
func (c Client) WeeklyTimeSeries(ctx context.Context, symbol string) (WeeklyTimeSeries, error) {
	url := fmt.Sprintf(urlTemplate, "TIME_SERIES_WEEKLY", c.URL, symbol, c.APIKey)

	data, err := c.get(ctx, url)
	if err != nil {
		return WeeklyTimeSeries{}, fmt.Errorf("could not get weekly time series: %w", err)
	}

	var wts WeeklyTimeSeries
	err = json.Unmarshal(data, &wts)
	if err != nil {
		return WeeklyTimeSeries{}, fmt.Errorf("could not unmarshal response: %w", err)
	}

	return wts, nil
}

func (c Client) get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create new request: %w", err)
	}

	resp, err := c.Doer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in request: %w", err)
	}
	defer resp.Body.Close() // nolint: errcheck

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return data, nil
}
