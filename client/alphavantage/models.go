package alphavantage

import (
	"strconv"

	"github.com/indiependente/go-stock/models"
)

// MetaData related to a quote request.
type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

// GlobalQuote represents the information about a global quote.
type GlobalQuote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

// Quote converts a GlobalQuote into a generic Quote.
func (gq GlobalQuote) Quote() models.Quote {
	q := models.Quote{}
	o, _ := strconv.ParseFloat(gq.Open, 64)
	q.Open = o
	h, _ := strconv.ParseFloat(gq.High, 64)
	q.High = h
	l, _ := strconv.ParseFloat(gq.Low, 64)
	q.Low = l
	p, _ := strconv.ParseFloat(gq.Price, 64)
	q.Price = p
	v, _ := strconv.ParseFloat(gq.Volume, 64)
	q.Volume = v

	return q
}

// TimeSeries is a Alpha Vantage time series.
type TimeSeries map[string]GlobalQuote

// TimeSeries represents weekly stock data for the quote over the last 20 years.
type WeeklyTimeSeries struct {
	MetaData MetaData   `json:"Meta Data"`
	Series   TimeSeries `json:"Weekly Time Series"`
}

// MonthlyTimeSeries represents monthly stock data for the quote over the last 20 years.
type MonthlyTimeSeries struct {
	MetaData MetaData   `json:"Meta Data"`
	Series   TimeSeries `json:"Monthly Time Series"`
}

// TimeSeries returns a generic quotes time series.
func (ts TimeSeries) TimeSeries() models.TimeSeries {
	t := models.TimeSeries{
		Labels: []string{},
		Points: []models.Quote{},
	}
	for k, v := range ts {
		t.Labels = append(t.Labels, k)
		t.Points = append(t.Points, v.Quote())
	}
	return t
}
