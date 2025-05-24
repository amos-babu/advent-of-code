package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/day1/input.txt")
	check(err)
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			continue
		}

		l, err1 := strconv.Atoi(parts[0])
		r, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		left = append(left, l)
		right = append(right, r)
	}

	rightFreq := make(map[int]int)
	for _, r := range right {
		rightFreq[r]++
	}

	totalScore := 0
	for _, l := range left {
		totalScore += l * rightFreq[l]
	}

	fmt.Println(totalScore)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
