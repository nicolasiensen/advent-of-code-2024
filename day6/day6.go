package main

import (
	"fmt"
	"os"
)

func main() {
	inputb, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width, height, obstacles, guard := Parse(string(inputb))
	guard.Patrol(width, height, obstacles)

	fmt.Printf("Width: %d Height: %d Unique positions: %v Obstacles: %d\n", width, height, len(guard.UniquePositions()), len(obstacles))
}
