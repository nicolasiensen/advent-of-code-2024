package main

import "strings"

func Parse(input string) (width, height int, obstacles []Point, guardPosition Point, guardBearing Bearing) {
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]
	width = len(lines[0])
	height = len(lines)

	for i, line := range lines {
		for j, col := range line {
			if col == '#' {
				obstacles = append(obstacles, Point{j, i})
			} else if col == '^' {
				guardBearing = North
				guardPosition = Point{j, i}
			} else if col == '>' {
				guardBearing = East
				guardPosition = Point{j, i}
			} else if col == 'v' {
				guardBearing = South
				guardPosition = Point{j, i}
			} else if col == '<' {
				guardBearing = West
				guardPosition = Point{j, i}
			}
		}
	}

	return width, height, obstacles, guardPosition, guardBearing
}
