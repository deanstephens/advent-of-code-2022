package utils

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (c Coordinate) Add(b Coordinate) Coordinate {
	return Coordinate{X: c.X + b.X, Y: c.Y + b.Y}
}
