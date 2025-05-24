package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input2.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := parseLineToIts(line)
		if isSafeReport(nums) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func removeUnsafe() {

}

func parseLineToIts(line string) []int {
	parts := strings.Fields(line)
	nums := make([]int, len(parts))

	for i, part := range parts {
		nums[i], _ = strconv.Atoi(part)
	}
	return nums
}

func isSafeReport(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	direction := 0

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if diff == 0 || abs(diff) > 3 {
			return false
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else {
			if (direction == 1 && diff < 0) || (direction == -1 && diff > 0) {
				return false
			}
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
