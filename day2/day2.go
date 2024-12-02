package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

			delta := level1 - level2

			if delta > 3 || delta < -3 || delta == 0 {
				break
			}

			if i == 0 {
				reportIncreasing = level1 < level2
			} else {
				if level1 < level2 && !reportIncreasing {
					break
				} else if level1 > level2 && reportIncreasing {
					break
				}
			}
		}
	}

	println("Safe reports:", safeReportsCount)
}
