package main

import (
	"fmt"
	"time"

	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
	"github.com/billwang7599/BillGame/internal/systems"
	"github.com/gdamore/tcell/v2"
)

func NewEntity(world *ecs.World, x int, y int, s rune) {
	entity := world.NewEntity()
	world.AddComponent(entity, components.NewMove(components.Right, 1, true))
	world.AddComponent(entity, components.NewPosition(x, y))
	world.AddComponent(entity, components.NewSprite(s))
}

func main() {
	fmt.Println("hello im client")
	screen, err := tcell.NewScreen()
	if err != nil {
		panic("Can't make new screen")
	}
	if err := screen.Init(); err != nil {
		panic("Screen failed to initialize")
	}
	defer screen.Fini()
	// creating world
	world := ecs.NewWorld()
	world.AddSystem(systems.NewInputSystem(world, screen))
	world.AddSystem(systems.NewMovementSystem(world))
	world.AddSystem(systems.NewRenderSystem(world, screen))
	NewEntity(world, 1, 1, 'X')
	NewEntity(world, 2, 1, 'Y')
	NewEntity(world, 1, 2, 'Z')
	entity := world.NewEntity()
	world.AddComponent(entity, components.NewMove(components.None, 1, false))
	world.AddComponent(entity, components.NewPosition(10, 10))
	world.AddComponent(entity, components.NewSprite('B'))

	const targetFrameTime = time.Second / 60

	// Create a new ticker that "ticks" every 16.67ms.
	ticker := time.NewTicker(targetFrameTime)
	defer ticker.Stop() // Clean up the ticker when we're done.

	// The main loop.
	for range ticker.C {
		// The loop will automatically run once per tick (approx. 60 times per second).
		// Put your repeating logic here.
		if !world.Running {
			break
		}

		world.Update()
	}
}
