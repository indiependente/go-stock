package models

// Quote represents a generic quote data.
type Quote struct {
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

// TimeSeries represents a time series.
type TimeSeries struct {
	Labels []string
	Points []Quote
}

// Price returns the price data points of the time series.
func (ts TimeSeries) Price() []float64 {
	data := []float64{}
	for _, q := range ts.Points {
		data = append(data, q.Price)
	}
	return data
}