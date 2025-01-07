package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputb, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input := string(inputb)
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]
	width := len(lines[0])
	height := len(lines)
	obstacles := []Point{}
	var position Point
	var bearing Bearing

	for i, line := range lines {
		for j, col := range line {
			if col == '#' {
				obstacles = append(obstacles, Point{j, i})
			} else if col == '^' {
				bearing = North
				position = Point{j, i}
			} else if col == '>' {
				bearing = East
				position = Point{j, i}
			} else if col == 'v' {
				bearing = South
				position = Point{j, i}
			} else if col == '<' {
				bearing = West
				position = Point{j, i}
			}
		}
	}

	guard := Guard{[]Point{position}, bearing}

	for guard.IsWithinBounds(width, height) {
		nextPos := guard.NextPos()

		for _, obstacle := range obstacles {
			if obstacle == nextPos {
				guard.Turn()
				break
			}
		}

		guard.Move()
	}

	fmt.Printf("Width: %d Height: %d Unique positions: %v Obstacles: %d\n", width, height, len(guard.UniquePositions()), len(obstacles))
}
