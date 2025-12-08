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
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var freshIdRanges []string
	var AvailableIds []string
	inFirst := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inFirst = false
			continue
		}
		if inFirst {

			freshIdRanges = append(freshIdRanges, line)
		} else {

			AvailableIds = append(AvailableIds, line)
		}
	}

	count := 0
	seen := make(map[int]bool)

	for _, j := range AvailableIds {
		aID, err := strconv.Atoi(j)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range freshIdRanges {
			parts := strings.Split(v, "-")
			a, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}

			b, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}

			if aID >= a && aID <= b && !seen[aID] {
				count++
				seen[aID] = true
			}

		}
	}

	fmt.Println(count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
