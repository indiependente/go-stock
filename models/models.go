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

func (ts TimeSeries) Len() int {
	return len(ts.Labels)
}
func (ts TimeSeries) Less(i, j int) bool {
	return ts.Labels[i] < ts.Labels[j]
}
func (ts TimeSeries) Swap(i, j int) {
	ts.Labels[i], ts.Labels[j] = ts.Labels[j], ts.Labels[i]
	ts.Points[i], ts.Points[j] = ts.Points[j], ts.Points[i]
}

// High returns the high data points of the time series.
func (ts TimeSeries) High() []float64 {
	data := []float64{}
	for _, q := range ts.Points {
		data = append(data, q.High)
	}
	return data
}
