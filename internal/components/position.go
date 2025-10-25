package components

import "fmt"

type Position struct {
	X int64
	Y int64
}

func NewPosition(x, y int64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("Position (%d, %d)", p.X, p.Y)
}
