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

	var pos = 50
	var positionNumbers []int
	var zeroCount int

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[:1]
		rotation, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatalf("Bad Number: %v", err)
		}
		if dir == "R" {
			for i := 0; i < rotation; i++ {
				pos = (pos + 1) % 100
				if pos == 0 {
					zeroCount++
				}
			}
		} else {
			for i := 0; i < rotation; i++ {
				pos--
				if pos < 0 {
					pos += 100
				}
				if pos == 0 {
					zeroCount++
				}
			}
		}

		positionNumbers = append(positionNumbers, pos)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	fmt.Println(zeroCount)
}
