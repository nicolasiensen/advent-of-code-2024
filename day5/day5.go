package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fstring := string(f)

	orderRegexp, err := regexp.Compile(`\d+\|\d+`)
	if err != nil {
		fmt.Println(err)
	}
	orders := orderRegexp.FindAllString(fstring, -1)

	updateRegexp, err := regexp.Compile(`(\d+\,)+(\d+)`)
	if err != nil {
		fmt.Println(err)
	}
	updates := updateRegexp.FindAllString(fstring, -1)

	sum := 0

	for _, update := range updates {
		if Correct(update, orders) {
			midv, err := MidValue(update)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			midi, err := strconv.Atoi(midv)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			sum += midi
		}
	}

	fmt.Println(sum)
}

func MidValue(update string) (string, error) {
	uarray := strings.Split(update, ",")

	if len(uarray)%2 == 0 {
		return "", error(nil)
	}

	item := uarray[(len(uarray) / 2)]

	return item, nil
}

func Correct(update string, orders []string) bool {
	for _, order := range orders {
		oarray := strings.Split(order, "|")
		o1index := strings.Index(update, oarray[0])
		o2index := strings.Index(update, oarray[1])

		if o1index != -1 && o2index != -1 && o1index > o2index {
			return false
		}
	}

	return true
}
