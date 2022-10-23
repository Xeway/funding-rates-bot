package models

type FundingRates struct {
	Success bool     `json:"success"`
	Result  []Result `json:"result"`
}

type Result struct {
	Future string  `json:"future"`
	Rate   float64 `json:"rate"`
	Time   string  `json:"time"`
}
