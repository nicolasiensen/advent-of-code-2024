package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func processInput(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		num1, err := strconv.Atoi(words[0])
		list1 = append(list1, num1)
		if err != nil {
			return nil, nil, err
		}

		num2, err := strconv.Atoi(words[len(words)-1])
		list2 = append(list2, num2)
		if err != nil {
			return nil, nil, err
		}
	}

	return list1, list2, nil
}

func calcDistances(list1 []int, list2 []int) []int {
	slices.Sort(list1)
	slices.Sort(list2)

	var distances []int

	for i := range list1 {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance = -distance
		}

		distances = append(distances, distance)
	}

	return distances
}

func calcSimiliarities(list1 []int, list2 []int) []int {
	var similiarities []int

	for i := range list1 {
		similiarity := 0

		for j := range list2 {
			if list1[i] == list2[j] {
				similiarity++
			}
		}

		similiarities = append(similiarities, list1[i]*similiarity)
	}

	return similiarities
}

func sum(list []int) int {
	sum := 0

	for i := range list {
		sum += list[i]
	}

	return sum
}

func main() {
	list1, list2, err := processInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	distances := calcDistances(list1, list2)
	distance := sum(distances)

	similiarities := calcSimiliarities(list1, list2)
	similiarity := sum(similiarities)

	fmt.Println("distance:", distance, "similiarity:", similiarity)
}
