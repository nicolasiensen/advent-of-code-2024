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

	width, height, obstacles, guardPosition, guardBearing := Parse(string(inputb))
	guard := BuildGuard(guardPosition, guardBearing)
	guard.Patrol(width, height, obstacles)

	fmt.Printf("Width: %d Height: %d Unique positions: %v Obstacles: %d\n", width, height, len(guard.UniquePositions()), len(obstacles))

	uniquePos := guard.UniquePositions()
	loopCount := 0

	for i := range uniquePos {
		newObstacles := make([]Point, len(obstacles)+1)
		copy(newObstacles, obstacles)
		newObstacles = append(newObstacles, uniquePos[i])
		guard := BuildGuard(guardPosition, guardBearing)
		guard.Patrol(width, height, newObstacles)
		if guard.IsInLoop {
			loopCount++
		}

		fmt.Printf("Obstacle at %v generates a loop: %v\n", uniquePos[i], guard.IsInLoop)
	}

	fmt.Printf("Loop count: %d\n", loopCount)
}
