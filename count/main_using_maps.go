package main

import (
	"fmt"
	"log"

	"github.com/himanshudogra/datafile"
)

func main() {
	lines, err := datafile.GetString("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("%s: %d\n", name, count)
	}
}
