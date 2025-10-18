package components

import "fmt"

type position struct {
	X int
	Y int
}

func NewPosition(x, y int) *position {
	return &position{
		X: x,
		Y: y,
	}
}

func (p *position) String() string {
	return fmt.Sprintf("Position (%d, %d)", p.X, p.Y)
}
