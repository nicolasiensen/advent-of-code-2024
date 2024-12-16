package main

import "testing"

var orders []string

func init() {
	orders = []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}
}

func TestCorrectWhenTrue(t *testing.T) {
	update := "75,47,61,53,29"

	correct := Correct(update, orders)

	expect := true

	if correct != expect {
		t.Fatalf(`Correct(%v, %v) = %v but it should be %v`, update, orders, correct, expect)
	}
}

func TestCorrectWhenFalse(t *testing.T) {
	update := "75,97,47,61,53"

	correct := Correct(update, orders)

	expect := false

	if correct != expect {
		t.Fatalf(`Correct(%v, %v) = %v but it should be %v`, update, orders, correct, expect)
	}
}

func TestMidValueWithEvenUpdate(t *testing.T) {
	update := "75,97,47,61"

	midv, err := MidValue(update)
	if err != nil {
		t.Fatalf(`MidValue(%v) = %v, %v but it should be %v, %v`, update, midv, err, nil, error(nil))
	}
}

func TestMidValueWithValidUpdate(t *testing.T) {
	update := "75,97,47,61,53"

	expected := "47"

	midv, err := MidValue(update)
	if midv != expected {
		t.Fatalf(`MidValue(%v) = %v, %v but it should be %v, %v`, update, midv, err, expected, nil)
	}
}
