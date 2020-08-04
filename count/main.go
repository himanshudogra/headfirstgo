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

	fmt.Println(lines)
}
