package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/Xeway/funding-rates-bot/models"
)

const FTX_API_URL = "https://ftx.com/api/funding_rates"

func main() {
	fundingRates := FetchFundingRatesAPI(FTX_API_URL)

	bestFundingRate := FindBestOpportunity(fundingRates.Result)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	FTX_API_KEY := os.Getenv("FTX_API_KEY")
	FTX_API_SECRET := os.Getenv("FTX_API_SECRET")

	fmt.Println(bestFundingRate)
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
