package models

type Result struct {
	Symbol               string  `json:"symbol"`
	MarkPrice            float64 `json:"markPrice"`
	IndexPrice           float64 `json:"indexPrice"`
	EstimatedSettlePrice float64 `json:"estimatedSettlePrice"`
	LastFundingRate      float64 `json:"lastFundingRate"`
	InterestRate         float64 `json:"interestRate"`
	NextFundingTime      int     `json:"nextFundingTime"`
	Time                 int     `json:"time"`
}

type FundingRates []Result
