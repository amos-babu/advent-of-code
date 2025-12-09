package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(grid) < 2 {
		log.Fatal("Not enough lines")
	}

	numRows := len(grid) - 1
	numCols := len(grid[0])
	opRow := grid[len(grid)-1]

	// Step 1: detect problem boundaries
	problemIndices := []int{}
	for c := 0; c < numCols; c++ {
		empty := true
		for r := 0; r < numRows; r++ {
			if c < len(grid[r]) && grid[r][c] != ' ' {
				empty = false
				break
			}
		}
		if empty {
			problemIndices = append(problemIndices, c)
		}
	}
	problemIndices = append(problemIndices, numCols) // add end

	// Step 2: extract problems
	var problems [][]int
	var operators []rune

	start := 0
	for _, end := range problemIndices {
		if end > start {
			// collect numbers
			var numbers []int
			for c := start; c < end; c++ {
				// check if this column is all spaces (skip)
				allSpace := true
				for r := 0; r < numRows; r++ {
					if c < len(grid[r]) && grid[r][c] != ' ' {
						allSpace = false
						break
					}
				}
				if allSpace {
					continue
				}
				// build number from top to bottom
				numStr := ""
				for r := 0; r < numRows; r++ {
					if c < len(grid[r]) && grid[r][c] != ' ' {
						numStr += string(grid[r][c])
					}
				}
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
			if len(numbers) > 0 {
				problems = append(problems, numbers)
				// operator is bottom-most non-space in this range
				op := ' '
				for c2 := start; c2 < end; c2++ {
					if c2 < len(opRow) && opRow[c2] != ' ' {
						op = opRow[c2]
						break
					}
				}
				operators = append(operators, op)
			}
		}
		start = end + 1
	}

	// Step 3: compute grand total, right-to-left
	grandTotal := 0
	for i := len(problems) - 1; i >= 0; i-- {
		numbers := problems[i]
		op := operators[i]
		result := 0
		if op == '*' {
			result = 1
			for _, n := range numbers {
				result *= n
			}
		} else if op == '+' {
			for _, n := range numbers {
				result += n
			}
		}
		grandTotal += result
	}

	fmt.Println("Grand total:", grandTotal)
}
