package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(file), "\n")
	count := 0

	for i := 0; i < len(lines)-3; i++ {
		line := lines[i]

		for j := 0; j < len(line)-2; j++ {
			tl := string(line[j])
			tr := string(line[j+2])
			md := string(lines[i+1][j+1])
			bl := string(lines[i+2][j])
			br := string(lines[i+2][j+2])
			x1 := tl + md + br
			x2 := bl + md + tr

			if (x1 == "MAS" || x1 == "SAM") && (x2 == "MAS" || x2 == "SAM") {
				count++
			}
		}
	}

	fmt.Println(count)
}
