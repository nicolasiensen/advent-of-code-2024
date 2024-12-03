package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"example.com/report"
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
		reportStr := strings.Split(line, " ")
		var reportInt []int

		for i := range reportStr {
			item, err := strconv.Atoi(reportStr[i])
			if err != nil {
				println(err)
			}

			reportInt = append(reportInt, item)
		}

		if report.Safe(reportInt) {
			safeReportsCount++
		}
	}

	println("Safe reports:", safeReportsCount)
}
