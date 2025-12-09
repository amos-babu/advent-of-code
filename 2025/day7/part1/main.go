package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Initialize the beam queue with that position
type Beam struct {
	row, column int
}

func main() {
	f, err := os.Open("./sample-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	beam := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		beam = append(beam, line)
	}

	matrix := make([][]rune, len(beam))

	for i, v := range beam {
		matrix[i] = []rune(v)
	}

	// for _, row := range matrix {
	// 	for _, cell := range row {
	// 		fmt.Printf("%c", cell)
	// 	}
	// 	fmt.Println()
	// }

	// Find the S position
	var startR, startC int
	found := false

	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 'S' {
				startR, startC = r, c
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	queue := []Beam{{startR, startC}}
	fmt.Println(queue)

	// Make a visited matrix

	// Use the simulation loop I described earlier

	// Count the splitters (^) where you actually split

}
