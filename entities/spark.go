package entities

type Timestamp struct {
	Timestamp []uint64 `json:"timestamp"`
}

type Close struct {
	Close []float64 `json:"close"`
}

type Spark struct {
	DataGranularity    float64   `json:"dataGranularity"`
	PreviousClose      float64   `json:"previousClose"`
	Symbol             string    `json:"symbol"`
	ChartPreviousClose float64   `json:"chartPreviousClose"`
	Timestamp          []int64   `json:"timestamp"`
	End                string    `json:"end"`
	Start              string    `json:"start"`
	Close              []float64 `json:"close"`
}

type SparkQuery struct {
	Symbol   string
	Interval string
	Range    string
}

type Quotes struct {
	Symbol []Spark `json:"symbol"`
}
