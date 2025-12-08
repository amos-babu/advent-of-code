package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rows := len(grid)
	if rows == 0 {
		return
	}

	cols := len(grid[0])

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*self*/, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	accessible := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}

			count := 0

			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					if grid[nr][nc] == '@' {
						count++
					}
				}
			}
			if count < 4 {
				accessible++
			}
		}
	}

	fmt.Println(accessible)

}
