package main

import (
	"fmt"
	"log"
	"math"

	datafile "github.com/himanshudogra/datafile/slices/float64"
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
	var err error
	numbers, err = datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(maximum(numbers...))
}
