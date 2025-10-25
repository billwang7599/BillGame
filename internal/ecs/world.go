package ecs

import (
	"net"
	"reflect"
)

// World is the central orchestrator, containing all the managers.
type World struct {
	entities    *EntityManager
	components  *ComponentManager
	systems     *SystemManager
	Clients     map[string]Entity // map client ip string to their entity
	Running     bool
	PacketQueue []any
	Conn        *net.UDPConn // shared UDP connection for send/receive
}

// NewWorld creates an initialized World with all its managers ready to go.
// Takes currentPort and serverAddr (ip:port) as strings.
func NewWorld(currentPort string, serverAddr string) *World {
	addr, err := net.ResolveUDPAddr("udp", ":"+currentPort)
	if err != nil {
		panic("Error resolving UDP address: " + err.Error())
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic("Error listening: " + err.Error())
	}

	world := &World{
		entities:    NewEntityManager(),
		components:  NewComponentManager(),
		systems:     NewSystemManager(),
		Clients:     make(map[string]Entity),
		PacketQueue: make([]any, 0),
		Running:     true,
		Conn:        conn,
	}
	// Add server address to Clients map
	world.Clients[serverAddr] = 0
	return world
}

func (w *World) AddToPacketQueue(packet any) {
	w.PacketQueue = append(w.PacketQueue, packet)
}

// These methods provide a clean API and hide the internal managers.
func (w *World) Print() {
	w.components.PrintComponent()
}
func (w *World) NewEntity() Entity {
	return w.entities.NewEntity()
}

func (w *World) AddComponent(entity Entity, component any) {
	if reflect.TypeOf(component).Kind() != reflect.Pointer {
		panic("Component must be a pointer")
	}
	w.components.AddComponent(entity, component)
}

func (w *World) AddSystem(system System) {
	w.systems.Add(system)
}

func (w *World) GetEntitiesOfComponent(component any) map[Entity]any {
	return w.components.GetComponent(component)
}

// You would add other facade methods here for getting, removing, and querying components.
// For example:
// func (w *World) GetComponent(...)
// func (w *World) Query(...)

// --- Game Loop Method ---

// Update runs one tick of the game loop, updating all systems.
func (w *World) Update() {
	w.systems.Update()
}
