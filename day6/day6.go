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
	uniquePositions := []Point{}
	var position Point
	var bearing Bearing

	for i, line := range lines {
		for j, col := range line {
			if col == '#' {
				obstacles = append(obstacles, Point{j, i})
			} else if col == '^' {
				bearing = north
				position = Point{j, i}
			} else if col == '>' {
				bearing = east
				position = Point{j, i}
			} else if col == 'v' {
				bearing = south
				position = Point{j, i}
			} else if col == '<' {
				bearing = west
				position = Point{j, i}
			}
		}
	}

	guard := Guard{[]Point{position}, bearing}

	for guard.Positions[len(guard.Positions)-1].X > 0 && guard.Positions[len(guard.Positions)-1].X < width && guard.Positions[len(guard.Positions)-1].Y > 0 && guard.Positions[len(guard.Positions)-1].Y < height {
		lastPos := guard.Positions[len(guard.Positions)-1]
		newPos := Point{lastPos.X, lastPos.Y}

		switch guard.Bearing {
		case north:
			newPos.Y -= 1
		case east:
			newPos.X += 1
		case south:
			newPos.Y += 1
		case west:
			newPos.X -= 1
		}

		isColliding := false

		for _, obstacle := range obstacles {
			if obstacle == newPos {
				isColliding = true
				break
			}
		}

		if isColliding {
			guard.Bearing += 1
			if guard.Bearing > 3 {
				guard.Bearing = 0
			}
		} else {
			guard.Positions = append(guard.Positions, newPos)
		}

		for i, p := range guard.Positions {
			unique := true

			for _, up := range uniquePositions {
				if p.X == up.X && p.Y == up.Y {
					unique = false
				}
			}

			if unique {
				fmt.Printf("Unique position: %v, %d\n", p, i)
				uniquePositions = append(uniquePositions, p)
			}
		}
	}

	fmt.Printf("Width: %d Height: %d Unique positions: %v Obstacles: %d\n", width, height, len(uniquePositions), len(obstacles))
}

type Point struct {
	X int
	Y int
}

type Guard struct {
	Positions []Point
	Bearing   Bearing
}

type Bearing int

const (
	north Bearing = iota
	east
	south
	west
)
