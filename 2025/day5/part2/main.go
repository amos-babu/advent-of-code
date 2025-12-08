package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var freshIdRanges []Range
	inFirst := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inFirst = false
			continue
		}
		if inFirst {
			parts := strings.Split(line, "-")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			freshIdRanges = append(freshIdRanges, Range{a, b})
		}
	}

	// Sort ranges by start
	sort.Slice(freshIdRanges, func(i, j int) bool {
		return freshIdRanges[i].start < freshIdRanges[j].start
	})

	// Merge overlapping ranges
	merged := []Range{}
	for _, r := range freshIdRanges {
		if len(merged) == 0 || merged[len(merged)-1].end < r.start-1 {
			merged = append(merged, r)
		} else {
			if r.end > merged[len(merged)-1].end {
				merged[len(merged)-1].end = r.end
			}
		}
	}

	// Count unique numbers
	count := 0
	for _, r := range merged {
		count += r.end - r.start + 1
	}

	fmt.Println("Unique count:", count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
