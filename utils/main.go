package utils

import (
	"log"
	"strconv"
)

func StringToFloat(x string) float64 {
	res, err := strconv.ParseFloat(x, 64)
	if err != nil {
		log.Fatal("Error for converting string to float64")
	}

	return res
}
