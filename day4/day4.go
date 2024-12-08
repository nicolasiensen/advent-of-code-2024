package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var rows []string
	var cols []string
	var diagL []string
	var diagR []string

	scanner := bufio.NewScanner(file)
	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		rows = append(rows, line)

		for i := 0; i < lineLen; i++ {
			charAsc := string(line[i])
			charDsc := string(line[lineLen-i-1])

			if len(cols) > i {
				cols[i] = cols[i] + charAsc
			} else {
				cols = append(cols, charAsc)
			}

			if len(diagL) > i+lineIndex {
				diagL[i+lineIndex] = diagL[i+lineIndex] + charAsc
			} else {
				diagL = append(diagL, charAsc)
			}

			if len(diagR) > i+lineIndex {
				diagR[i+lineIndex] = diagR[i+lineIndex] + charDsc
			} else {
				diagR = append(diagR, charDsc)
			}
		}

		lineIndex++
	}

	count := 0

	for _, text := range rows {
		count += strings.Count(text, "XMAS")
		count += strings.Count(text, "SAMX")
	}

	for _, text := range cols {
		count += strings.Count(text, "XMAS")
		count += strings.Count(text, "SAMX")
	}

	for _, text := range diagL {
		count += strings.Count(text, "XMAS")
		count += strings.Count(text, "SAMX")
	}

	for _, text := range diagR {
		count += strings.Count(text, "XMAS")
		count += strings.Count(text, "SAMX")
	}

	fmt.Println(count)
}
