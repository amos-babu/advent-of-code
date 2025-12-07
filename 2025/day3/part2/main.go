package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		var batteriesLine []int
		line := scanner.Text()
		for _, v := range line {
			b := int(v - '0')
			batteriesLine = append(batteriesLine, b)
		}

		numDigits := 12
		start := 0
		maxNum := ""

		for k := 0; k < numDigits; k++ {
			remaining := numDigits - k
			end := len(batteriesLine) - remaining

			bestDigit := -1
			bestIndex := start

			for i := start; i <= end; i++ {
				if batteriesLine[i] > bestDigit {
					bestDigit = batteriesLine[i]
					bestIndex = i
				}
			}

			maxNum += strconv.Itoa(bestDigit)
			start = bestIndex + 1
		}
		bank, err := strconv.Atoi(maxNum)
		if err != nil {
			log.Fatal(err)
		}

		total += bank
	}

	fmt.Println(total)

}
