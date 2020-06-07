package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var errs []error
	var sum float64
	var sampleCount float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		sum += number
		sampleCount++
	}

	totalErrs := len(errs)
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
	fmt.Printf("Total errors: %d ", totalErrs)
	fmt.Println(" Errors: ", errs)
}
