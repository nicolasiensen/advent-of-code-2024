package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"example.com/level"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		println(err)
	}

	scanner := bufio.NewScanner(file)
	safeReportsCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		levelsStr := strings.Split(line, " ")
		var increasing bool

		for i := range levelsStr {
			// The report is safe when this for loop reaches the last level without breaking it
			if i+1 == len(levelsStr) {
				safeReportsCount++
				break
			}

			level1, err := strconv.Atoi(levelsStr[i])
			if err != nil {
				println(err)
			}

			level2, err := strconv.Atoi(levelsStr[i+1])
			if err != nil {
				println(err)
			}

			if i == 0 {
				increasing = level1 < level2
			}

			safe := level.Safe(level1, level2, increasing)

			if !safe {
				break
			}
		}
	}

	println("Safe reports:", safeReportsCount)
}
