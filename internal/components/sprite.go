package components

import "fmt"

type Sprite struct {
	Char rune
}

func NewSprite(char rune) *Sprite {
	return &Sprite{
		Char: char,
	}
}

func (s *Sprite) String() string {
	return fmt.Sprintf("Char %c", s.Char)
}
