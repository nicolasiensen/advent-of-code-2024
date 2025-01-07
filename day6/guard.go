package main

import (
	"fmt"
)

type Guard struct {
	Positions []Point
	Bearing   Bearing
}

func (g Guard) LastPos() Point {
	return g.Positions[len(g.Positions)-1]
}

func (g Guard) IsWithinBounds(width int, height int) bool {
	lastPos := g.LastPos()
	return lastPos.X > 0 && lastPos.X < width && lastPos.Y > 0 && lastPos.Y < height
}

func (g Guard) UniquePositions() []Point {
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

func (g *Guard) Move() {
	g.Positions = append(g.Positions, g.NextPos())
}

func (g Guard) NextPos() Point {
	nextPos := Point{g.LastPos().X, g.LastPos().Y}

	switch g.Bearing {
	case North:
		nextPos.Y -= 1
	case East:
		nextPos.X += 1
	case South:
		nextPos.Y += 1
	case West:
		nextPos.X -= 1
	}

	return nextPos
}

func (g *Guard) Turn() Bearing {
	g.Bearing += 1
	if g.Bearing > 3 {
		g.Bearing = 0
	}
	return g.Bearing
}
