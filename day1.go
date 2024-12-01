package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var list1 []int
	var list2 []int
	var distances []int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		num1, err := strconv.Atoi(words[0])
		list1 = append(list1, num1)
		if err != nil {
			fmt.Println(err)
		}

		num2, err := strconv.Atoi(words[len(words)-1])
		list2 = append(list2, num2)
		if err != nil {
			fmt.Println(err)
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sort.Slice(list2, func(i, j int) bool {
		return i < j
	})

	for i := range list1 {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance = -distance
		}

		distances = append(distances, distance)
	}

	sum := 0

	for i := range distances {
		sum += distances[i]
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
