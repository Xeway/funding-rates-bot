package utils

import (
	"log"
	"strconv"
)

// Convert the string in parameter into a float64 and return it
func StringToFloat(x string) float64 {
	res, err := strconv.ParseFloat(x, 64)
	if err != nil {
		log.Fatal("Error for converting string to float64")
	}

	return res
}
