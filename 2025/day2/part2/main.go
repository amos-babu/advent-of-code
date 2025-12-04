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
	f, err := os.Open("../sample-input.txt")
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

			if len(s) < 2 {
				continue
			}

			for i := 2; i < (len(s)/2)+1; i++ {
				n := s[:i]
				fmt.Println(n)
				if strings.Contains(s, n) {
					num, err := strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
					totalInvalid += num
				}
				break
			}
		}

		fmt.Println(totalInvalid)
		// break
	}
}

func dashSeparator(r rune) bool {
	return r == '-'
}

func separator(r rune) bool {
	return r == ','
}
