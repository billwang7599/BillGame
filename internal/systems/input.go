package systems

import (
	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
	"github.com/gdamore/tcell/v2"
)

type InputSystem struct {
	world     *ecs.World
	inputChan chan rune
	screen    tcell.Screen
}

func NewInputSystem(w *ecs.World, screen tcell.Screen) *InputSystem {
	ch := make(chan rune)
	is := &InputSystem{
		world:     w,
		inputChan: ch,
		screen:    screen,
	}
	go is.PollInput(ch)
	return is
}

func (is *InputSystem) PollInput(ch chan rune) {
	for {
		ev := is.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune {
				ch <- ev.Rune()
			}

		default:
		}

	}
}

func (is *InputSystem) Update() {
	select {
	case input, ok := <-is.inputChan:
		if !ok {
			return
		}
		// maybe add input to some sort of queue system
		// right now just gets first char of input
		if input == 'q' {
			is.world.Running = false
		}
		is.MovementUpdate(input)
	default:
		// No input available
	}
}

func (is *InputSystem) MovementUpdate(key rune) {
	entityMapping := is.world.GetEntitiesOfComponent(components.Move{})
	dir := components.None
	switch key {
	case 'w':
		dir = components.Up
	case 'a':
		dir = components.Left
	case 's':
		dir = components.Down
	case 'd':
		dir = components.Right
	default:
	}

	for _, move := range entityMapping {
		move := move.(*components.Move)
		move.Direction = dir
	}
}
