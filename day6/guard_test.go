package main

import (
	"os"
	"strconv"
	"testing"
)

func TestPatrolWithoutLoop(t *testing.T) {
	inputb, _ := os.ReadFile("./fixtures/input_without_loop.txt")
	width, height, obstacles, guard := Parse(string(inputb))

	guard.Patrol(width, height, obstacles)
	expected := false
	result := guard.IsInLoop

	if expected != result {
		t.Fatalf(`IsInLoop = %v but expected %v`, result, expected)
	}
}

func TestPatrolWithLoop(t *testing.T) {
	for _, i := range []int{1, 2, 3, 4, 5, 6} {
		inputb, _ := os.ReadFile("./fixtures/input_with_loop" + strconv.Itoa(i) + ".txt")
		width, height, obstacles, guard := Parse(string(inputb))

		guard.Patrol(width, height, obstacles)
		expected := true
		result := guard.IsInLoop

		if expected != result {
			t.Fatalf(`IsInLoop = %v but expected %v`, result, expected)
		}

	}
}
