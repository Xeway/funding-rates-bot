package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Xeway/funding-rates-bot/models"
)

const FTX_API_URL = "https://ftx.com/api/funding_rates"

func main() {
	fundingRates := FetchAPI(FTX_API_URL)

	fmt.Println(fundingRates)
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

	return fundingRates
}
