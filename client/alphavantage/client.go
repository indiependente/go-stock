package alphavantage

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	url := fmt.Sprintf("%s/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", c.URL, symbol, c.APIKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("could not create new request: %w", err)
	}

	resp, err := c.Doer.Do(req)
	if err != nil {
		return GlobalQuoteResponse{}, fmt.Errorf("error in request: %w", err)
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
