package main

import (
	"encoding/json"
	"github.com/Xeway/funding-rates-bot/models"
	"github.com/joho/godotenv"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const FTX_API_URL = "https://ftx.com/api/funding_rates"

func main() {
	fundingRates := FetchFundingRatesAPI(FTX_API_URL)

	bestFundingRate := FindBestOpportunity(fundingRates.Result)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PerformTrade(bestFundingRate)
}

func FetchFundingRatesAPI(url string) models.FundingRates {
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

func PerformTrade(bestFundingRate models.FundingRates) {
	BINANCE_API_KEY := os.Getenv("BINANCE_API_KEY")
	BINANCE_API_SECRET := os.Getenv("BINANCE_API_SECRET")

	// perform trade with http post request
}
