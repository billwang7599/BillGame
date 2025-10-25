package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
	"github.com/billwang7599/BillGame/internal/systems"
	"github.com/gdamore/tcell/v2"
)

func main() {
	port := flag.String("port", "3002", "UDP port to send and receive packets on")
	serverAddr := flag.String("server-addr", "localhost:3000", "Server address ip:port")
	flag.Parse()
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

	world := ecs.NewWorld(*port, *serverAddr)
	world.AddSystem(systems.NewReceiveSystem(world.Conn, world)) // use shared UDPConn
	world.AddSystem(systems.NewInputSystem(world, screen))
	world.AddSystem(systems.NewMovementSystem(world))
	world.AddSystem(systems.NewRenderSystem(world, screen))
	world.AddSystem(systems.NewSendSystem(world, *serverAddr)) // use shared UDPConn and server address

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
