package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./sample-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
