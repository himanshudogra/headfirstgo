package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {

	var numbers []float64
	arguments := os.Args[1:]
	for _, argument := range arguments {
		argumentsNew, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, argumentsNew)
	}
	fmt.Println(maximum(numbers...))
}
