package systems

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
)

type NetworkReceiveSystem struct {
	queue [][]byte
	world *ecs.World
}

func NewReceiveSystem(world *ecs.World) *NetworkReceiveSystem {
	system := NetworkReceiveSystem{
		queue: make([][]byte, 0),
		world: world,
	}
	go system.handleConnection()
	return &system
}

func (nrs *NetworkReceiveSystem) handleConnection() {
	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := nrs.world.Conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}
		// track clients to know where to send after
		clientKey := clientAddr.String()
		if _, ok := nrs.world.Clients[clientKey]; !ok {
			nrs.world.Clients[clientKey] = 0 // placeholder entity
			// log.Printf("New client detected: %v\n", clientAddr)
		}
		// Add the received bytes to the queue
		packet := make([]byte, n)
		copy(packet, buffer[:n])
		// log.Printf("Received packet from %v: %x\n", clientAddr, packet)
		nrs.queue = append(nrs.queue, packet)
	}
}

func Deserialize[T any](reader *bytes.Reader) (T, error) {
	var packet T
	err := binary.Read(reader, binary.BigEndian, &packet)
	return packet, err
}

func (nrs *NetworkReceiveSystem) Update() {
	for _, packetBuffer := range nrs.queue {
		packetType := packetBuffer[0]
		reader := bytes.NewReader(packetBuffer)
		switch packetType {
		case ecs.PlayerInputPacketType:
			nrs.handlePlayerInputPacket(reader)
		case ecs.EntityPositionPacketType:
			nrs.handleEntityPositionPacket(reader)
		default:
			log.Printf("Received unknown packet type: %d\n", packetType)
		}
	}
	// Clear the queue after processing
	nrs.queue = nrs.queue[:0]
}

func (nrs *NetworkReceiveSystem) handlePlayerInputPacket(reader *bytes.Reader) {
	packet, err := Deserialize[ecs.PlayerInputPacket](reader)
	if err != nil {
		return
	}
	entity := packet.EntityId
	key := packet.Key
	entityComponentMapping := nrs.world.GetEntitiesOfComponent(components.Move{})
	moveComponent, ok := entityComponentMapping[entity]
	if !ok {
		return
	}
	move := moveComponent.(*components.Move)
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
	move.Direction = dir
}

func (nrs *NetworkReceiveSystem) handleEntityPositionPacket(reader *bytes.Reader) {
	packet, err := Deserialize[ecs.EntityPositionPacket](reader)
	if err != nil {
		return
	}
	entity := packet.EntityId
	newPositionComponent := packet.State
	entityComponentMapping := nrs.world.GetEntitiesOfComponent(components.Position{})
	positionComponent, ok := entityComponentMapping[entity]
	if !ok {
		return
	}
	position := positionComponent.(*components.Position)
	*position = newPositionComponent
}
