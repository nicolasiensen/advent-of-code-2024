package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func areAdjacentLevelsSafe(level1 int, level2 int, isFirstLevel bool, reportIncreasing bool) bool {
	delta := level1 - level2

	if delta > 3 || delta < -3 || delta == 0 {
		return false
	}

	if level1 < level2 && !reportIncreasing && !isFirstLevel {
		return false
	} else if level1 > level2 && reportIncreasing && !isFirstLevel {
		return false
	}
	return true
}

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
		var reportIncreasing bool

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
				reportIncreasing = level1 < level2
			}

			safe := areAdjacentLevelsSafe(level1, level2, i == 0, reportIncreasing)

			if !safe {
				break
			}
		}
	}

	println("Safe reports:", safeReportsCount)
}
