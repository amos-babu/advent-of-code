package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	enabled := true

	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`\bdo\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	for scanner.Scan() {
		line := scanner.Text()

		doLocs := doRe.FindAllStringIndex(line, -1)
		dontLocs := dontRe.FindAllStringIndex(line, -1)

		type toggle struct {
			pos    int
			enable bool
		}
		var toggles []toggle
		for _, loc := range doLocs {
			toggles = append(toggles, toggle{loc[0], true})
		}
		for _, loc := range dontLocs {
			toggles = append(toggles, toggle{loc[0], false})
		}

		for i := 0; i < len(toggles)-1; i++ {
			for j := i + 1; j < len(toggles); j++ {
				if toggles[i].pos > toggles[j].pos {
					toggles[i], toggles[j] = toggles[j], toggles[i]
				}
			}
		}

		mulMatches := mulRe.FindAllStringSubmatchIndex(line, -1)
		currToggle := 0

		for _, match := range mulMatches {
			start := match[0]

			for currToggle < len(toggles) && toggles[currToggle].pos <= start {
				enabled = toggles[currToggle].enable
				currToggle++
			}

			if !enabled {
				continue
			}

			num1Str := line[match[2]:match[3]]
			num2Str := line[match[4]:match[5]]
			num1, err1 := strconv.Atoi(num1Str)
			num2, err2 := strconv.Atoi(num2Str)
			checkErr(err1)
			checkErr(err2)

			total += num1 * num2
		}
	}

	fmt.Println("Part Two total:", total)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
