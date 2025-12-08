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
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))

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
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalAccessible := 0
	changed := true

	for changed {
		changed = false

		// mark positions to remove
		toRemove := make([][]bool, rows)
		for i := range toRemove {
			toRemove[i] = make([]bool, cols)
		}

		// count which @ become exposed this round
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
					toRemove[r][c] = true
				}
			}
		}

		// apply removals + update total count
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if toRemove[r][c] {
					grid[r][c] = '.'
					totalAccessible++
					changed = true
				}
			}
		}
	}

	fmt.Println("TOTAL accessible rolls:", totalAccessible)

}
