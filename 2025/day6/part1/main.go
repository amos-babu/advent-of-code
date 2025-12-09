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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line) // splits on spaces
		lines = append(lines, fields)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(lines) < 2 {
		log.Fatal("Not enough lines")
	}

	numRows := len(lines) - 1 // last line is operators
	numCols := len(lines[0])

	// Build problems column-wise
	problems := [][]string{}
	ops := lines[len(lines)-1] // last row

	for col := 0; col < numCols; col++ {
		problem := []string{}
		allEmpty := true
		for row := 0; row < numRows; row++ {
			if col < len(lines[row]) {
				val := lines[row][col]
				problem = append(problem, val)
				if val != "" {
					allEmpty = false
				}
			}
		}
		if !allEmpty {
			problem = append(problem, ops[col]) // add operator
			problems = append(problems, problem)
		}
	}

	// Solve each problem
	grandTotal := 0
	for _, p := range problems {
		op := p[len(p)-1]
		numbers := p[:len(p)-1]

		result := 0
		if op == "*" {
			result = 1
			for _, n := range numbers {
				num, _ := strconv.Atoi(n)
				result *= num
			}
		} else if op == "+" {
			result = 0
			for _, n := range numbers {
				num, _ := strconv.Atoi(n)
				result += num
			}
		}
		grandTotal += result
	}

	fmt.Println("Grand total:", grandTotal)
}
