package main

import (
	"fmt"
)

func inRange(min int, max int, numbers ...int) []int {

	var result []int
	for _, number := range numbers {
		if number > min && number < max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, 23, 46, -23, -99))
}
