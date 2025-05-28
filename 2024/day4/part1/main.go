package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("demo.txt")
	checkErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(string(grid[0][1]))
}

func checkErr(err error) {
	if err != nil {
		panic(err)

	}
}
