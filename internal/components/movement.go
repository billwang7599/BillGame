package components

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	None
)

type Move struct {
	Direction Direction
}

func NewMove(direction Direction) *Move {
	return &Move{
		Direction: direction,
	}
}
