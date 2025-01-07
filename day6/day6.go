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

	for guard.isWithinBounds(width, height) {
		newPos := Point{guard.lastPos().X, guard.lastPos().Y}

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
	}

	fmt.Printf("Width: %d Height: %d Unique positions: %v Obstacles: %d\n", width, height, len(guard.uniquePositions()), len(obstacles))
}

type Point struct {
	X int
	Y int
}

type Guard struct {
	Positions []Point
	Bearing   Bearing
}

func (g Guard) lastPos() Point {
	return g.Positions[len(g.Positions)-1]
}

func (g Guard) isWithinBounds(width int, height int) bool {
	return g.lastPos().X > 0 && g.lastPos().X < width && g.lastPos().Y > 0 && g.lastPos().Y < height
}

func (g Guard) uniquePositions() []Point {
	var uniquePositions []Point
	for i, p := range g.Positions {
		unique := true

		for _, up := range uniquePositions {
			if p == up {
				unique = false
			}
		}

		if unique {
			fmt.Printf("Unique position: %v, %d\n", p, i)
			uniquePositions = append(uniquePositions, p)
		}
	}
	return uniquePositions
}

func (g *Guard) move() {
	g.Positions = append(g.Positions, g.nextPos())
}

func (g Guard) nextPos() Point {
	nextPos := Point{g.lastPos().X, g.lastPos().Y}

	switch g.Bearing {
	case north:
		nextPos.Y -= 1
	case east:
		nextPos.X += 1
	case south:
		nextPos.Y += 1
	case west:
		nextPos.X -= 1
	}

	return nextPos
}

func (g *Guard) turn() Bearing {
	g.Bearing += 1
	if g.Bearing > 3 {
		g.Bearing = 0
	}
	return g.Bearing
}

type Bearing int

const (
	north Bearing = iota
	east
	south
	west
)
