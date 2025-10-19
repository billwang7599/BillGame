package components

import "fmt"

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("Position (%d, %d)", p.X, p.Y)
}
