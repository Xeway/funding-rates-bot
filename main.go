package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"

	"github.com/Xeway/funding-rates-bot/models"
)

const FTX_API_URL = "https://ftx.com/api/funding_rates"

func main() {
	fundingRates := FetchAPI(FTX_API_URL)

	bestFundingRate := FindBestOpportunity(fundingRates.Result)

	fmt.Println(bestFundingRate)
}

func FetchAPI(url string) models.FundingRates {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var fundingRates models.FundingRates
	json.Unmarshal(responseData, &fundingRates)

	if !fundingRates.Success {
		log.Fatal("API error.")
	}

	return fundingRates
}

func FindBestOpportunity(fundingRates []models.Result) models.Result {
	best := 0

	for i := 1; i < len(fundingRates); i++ {
		if math.Abs(fundingRates[i].Rate) > math.Abs(fundingRates[best].Rate) {
			best = i
		}
	}

	return fundingRates[best]
}
