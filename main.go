package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://ftx.com/api/funding_rates")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
