package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		println(err)
	}

	mulRegexp, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	ops := mulRegexp.FindAllString(string(file), -1)

	digRegexp, _ := regexp.Compile(`\d{1,3}`)
	sum := 0
	enabled := true

	for i := range ops {
		op := ops[i]

		if op == "don't()" {
			enabled = false
		} else if op == "do()" {
			enabled = true
		} else if enabled {
			args := digRegexp.FindAllString(op, -1)

			op1, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
			}

			op2, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
			}

			sum += op1 * op2
		}
	}

	fmt.Println("Result:", sum)
}
