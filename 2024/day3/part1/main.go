package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, err1 := strconv.Atoi(match[1])
			checkErr(err1)
			num2, err2 := strconv.Atoi(match[2])
			checkErr(err2)
			totalMatch := num1 * num2

			total += totalMatch
		}

	}

	fmt.Println(total)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
