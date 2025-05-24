package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("./input.txt")
	check(err)

	defer data.Close()

	scanner := bufio.NewScanner(data)

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

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

	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]

		if diff < 0 {
			diff = -diff
		}

		total += diff
	}

	fmt.Println(total)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
