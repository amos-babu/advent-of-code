package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var numbers []string

	for scanner.Scan() {
		line := scanner.Text()
		u := strings.FieldsFunc(line, separator)
		for _, num := range u {
			numbers = append(numbers, num)
		}
	}

	totalInvalid := 0

	for _, value := range numbers {
		strs := strings.FieldsFunc(value, dashSeparator)
		ints := make([]int, len(strs))
		for i, s := range strs {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			ints[i] = n
		}

		for i := ints[0]; i <= ints[1]; i++ {
			s := strconv.Itoa(i)

			if isInvalidID(s) {
				totalInvalid += i
			}
		}

		fmt.Println(totalInvalid)
	}
}

func isInvalidID(s string) bool {
	length := len(s)

	for subLen := 1; subLen <= length/2; subLen++ {
		if length%subLen != 0 {
			continue
		}

		sub := s[:subLen]
		repeated := strings.Repeat(sub, length/subLen)

		if repeated == s {
			return true
		}
	}

	return false
}

func dashSeparator(r rune) bool {
	return r == '-'
}

func separator(r rune) bool {
	return r == ','
}
