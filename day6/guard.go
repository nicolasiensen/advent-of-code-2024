package main

type Guard struct {
	Vectors  []Vector
	IsInLoop bool
}

func BuildGuard(position Point, bearing Bearing) Guard {
	return Guard{[]Vector{{position, bearing}}, false}
}

func (g Guard) Position() Point {
	return g.Vectors[len(g.Vectors)-1].Point
}

func (g Guard) IsWithinBounds(width int, height int) bool {
	lastPos := g.Position()
	return lastPos.X > 0 && lastPos.X < width && lastPos.Y > 0 && lastPos.Y < height
}

func (g Guard) UniquePositions() []Point {
	var uniquePositions []Point
	for _, v := range g.Vectors {
		unique := true

		for _, up := range uniquePositions {
			if v.Point == up {
				unique = false
			}
		}

		if unique {
			uniquePositions = append(uniquePositions, v.Point)
		}
	}
	return uniquePositions
}

func (g *Guard) Move() {
	g.Vectors = append(g.Vectors, g.NextVector())
}

func (g Guard) NextVector() Vector {
	nextVec := Vector{g.Position(), g.Bearing()}

	switch g.Bearing() {
	case North:
		nextVec.Point.Y -= 1
	case East:
		nextVec.Point.X += 1
	case South:
		nextVec.Point.Y += 1
	case West:
		nextVec.Point.X -= 1
	}

	return nextVec
}

func (g *Guard) Turn() Bearing {
	bearing := g.Bearing()
	bearing += 1
	if bearing > 3 {
		bearing = 0
	}
	g.Vectors = append(g.Vectors, Vector{g.Position(), bearing})
	return bearing
}

func (g *Guard) Patrol(width int, height int, obstacles []Point) {
	for g.IsWithinBounds(width, height) && !g.IsInLoop {
		nextPos := g.NextVector().Point

		for _, obstacle := range obstacles {
			if obstacle == nextPos {
				g.Turn()
				break
			}
		}

		g.Move()
		g.CheckLoop()
	}
}

func (g Guard) Bearing() Bearing {
	return g.Vectors[len(g.Vectors)-1].Bearing
}

func (g *Guard) CheckLoop() {
	for i := 0; i < len(g.Vectors); i++ {
		for j := i + 1; j < len(g.Vectors); j++ {
			if g.Vectors[i] == g.Vectors[j] {
				g.IsInLoop = true
			}
		}
	}
}
