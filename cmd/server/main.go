package main

import (
	"fmt"
	"time"

	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
	"github.com/billwang7599/BillGame/internal/systems"
)

func main() {
	fmt.Println("hello im server")
	receivePort := "3000"
	serverAddr := "localhost:3003"
	world := ecs.NewWorld(receivePort, serverAddr)
	world.AddSystem(systems.NewReceiveSystem(world.Conn, world)) // use shared UDPConn for receive
	world.AddSystem(systems.NewMovementSystem(world))
	world.AddSystem(systems.NewSendStateSystem(world))
	world.AddSystem(systems.NewSendSystem(world, serverAddr)) // use shared UDPConn for send

	entity := world.NewEntity()
	world.AddComponent(entity, components.NewMove(components.None, 1, false))
	world.AddComponent(entity, components.NewPosition(10, 10))
	world.AddComponent(entity, components.NewSprite('B'))
	world.AddComponent(entity, components.NewControllable())

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
