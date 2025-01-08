package main

import "strings"

func Parse(input string) (width, height int, obstacles []Point, guard Guard) {
	lines := strings.Split(input, "\n")
	lines = lines[0 : len(lines)-1]
	width = len(lines[0])
	height = len(lines)
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

	guard = BuildGuard(position, bearing)

	return width, height, obstacles, guard
}
