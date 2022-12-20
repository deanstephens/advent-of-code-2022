package utils

import "math"

type Coordinate3d struct {
	X int
	Y int
	Z int
}

func (c Coordinate3d) IsAdjacent(c2 Coordinate3d) bool {
	distance := int(math.Abs(float64(c.X-c2.X))) + int(math.Abs(float64(c.Y-c2.Y))) + int(math.Abs(float64(c.Z-c2.Z)))
	return distance == 1
}

func (c Coordinate3d) Add(b Coordinate3d) Coordinate3d {
	return Coordinate3d{X: c.X + b.X, Y: c.Y + b.Y, Z: c.Z + b.Z}
}

func (c Coordinate3d) Equals(b Coordinate3d) bool {
	return c.X == b.X && c.Y == b.Y && c.Z == b.Z
}
