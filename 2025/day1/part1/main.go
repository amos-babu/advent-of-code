package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		pos             = 50
		positionNumbers []int
	)

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[:1]
		rotation, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("Bad number: %v", err)
		}

		if dir == "R" {
			pos = (pos + rotation) % 100
		} else {
			pos = (pos - rotation + 100) % 100
			if pos < 0 {
				pos += 100
			}
		}
		positionNumbers = append(positionNumbers, pos)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fP := make(map[int]int)
	maxFrequency := 0

	for _, v := range positionNumbers {
		if _, exists := fP[v]; !exists {
			fP[v] = 1
		} else {
			fP[v]++
			if fP[v] > maxFrequency {
				maxFrequency = fP[v]
			}
		}
	}

	fmt.Println(maxFrequency)

}
