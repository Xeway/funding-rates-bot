package main

import (
	"encoding/json"
	"github.com/Xeway/funding-rates-bot/models"
	"github.com/Xeway/funding-rates-bot/utils"
	"github.com/joho/godotenv"
	"io"
	"log"
	"math"
	"net/http"
)

const BINANCE_API_URL = "https://fapi.binance.com/fapi/v1/premiumIndex"

func main() {
	fundingRates := FetchFundingRatesAPI(BINANCE_API_URL)

	bestFundingRate := FindBestOpportunity(fundingRates)

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

	return fundingRates
}

func FindBestOpportunity(fundingRates []models.Result) models.Result {
	best := 0

	for i := 1; i < len(fundingRates); i++ {
		if math.Abs(utils.StringToFloat(fundingRates[i].LastFundingRate)) > math.Abs(utils.StringToFloat(fundingRates[best].LastFundingRate)) {
			best = i
		}
	}

	return fundingRates[best]
}

func PerformTrade(bestFundingRate models.Result) {
	// BINANCE_API_KEY := os.Getenv("BINANCE_API_KEY")
	// BINANCE_API_SECRET := os.Getenv("BINANCE_API_SECRET")

	// perform trade with http post request
}
