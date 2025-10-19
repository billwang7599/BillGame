package components

import "fmt"

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	None
)

type Move struct {
	Direction  Direction
	Speed      int
	Continuous bool
}

func NewMove(direction Direction, speed int, continuous bool) *Move {
	return &Move{
		Direction:  direction,
		Speed:      speed,
		Continuous: continuous,
	}
}

func (m *Move) String() string {
	return fmt.Sprintf("Direction: %d || Speed: %d", m.Direction, m.Speed)
}
