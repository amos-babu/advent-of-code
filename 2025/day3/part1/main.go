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

		max := batteriesLine[0]
		maxId := 0

		for i, v := range batteriesLine[:len(batteriesLine)-1] {
			if v > max {
				max = v
				maxId = i
			}

		}

		secondMax := batteriesLine[maxId+1]
		for _, v := range batteriesLine[maxId+1:] {

			if v > secondMax {
				secondMax = v

			}
		}

		bank, err := strconv.Atoi(strconv.Itoa(max) + strconv.Itoa(secondMax))
		if err != nil {
			log.Fatal(err)
		}

		total += bank

		// break
	}

	fmt.Println(total)

}
