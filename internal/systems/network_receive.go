package systems

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"

	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
)

type NetworkReceiveSystem struct {
	connection *net.UDPConn
	queue      [][]byte
	world      *ecs.World
}

func NewReceiveSystem(conn *net.UDPConn, world *ecs.World) *NetworkReceiveSystem {
	system := NetworkReceiveSystem{
		connection: conn,
		queue:      make([][]byte, 0),
		world:      world,
	}
	go system.handleConnection()
	return &system
}

func (nrs *NetworkReceiveSystem) handleConnection() {
	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := nrs.connection.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error reading from UDP:", err)
			continue
		}
		// track clients to know where to send after
		clientKey := clientAddr.String()
		if _, ok := nrs.world.Clients[clientKey]; !ok {
			nrs.world.Clients[clientKey] = 0 // placeholder entity
			log.Printf("New client detected: %v\n", clientAddr)
		}
		// Add the received bytes to the queue
		packet := make([]byte, n)
		copy(packet, buffer[:n])
		log.Printf("Received packet from %v: %x\n", clientAddr, packet)
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
		log.Println(packetType)
		switch packetType {
		case ecs.PlayerInputPacketType:
			packet, err := Deserialize[ecs.PlayerInputPacket](reader)
			if err != nil {
				continue
			}
			entity := packet.EntityId
			key := packet.Key
			entityComponentMapping := nrs.world.GetEntitiesOfComponent(components.Move{})
			moveComponent, ok := entityComponentMapping[entity]
			if !ok {
				continue
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
		case ecs.EntityPositionPacketType:
			packet, err := Deserialize[ecs.EntityPositionPacket](reader)
			if err != nil {
				continue
			}
			entity := packet.EntityId
			newPositionComponent := packet.State
			entityComponentMapping := nrs.world.GetEntitiesOfComponent(components.Position{})
			positionComponent, ok := entityComponentMapping[entity]
			if !ok {
				continue
			}
			position := positionComponent.(*components.Position)
			*position = newPositionComponent
		default:
			log.Printf("Received unknown packet type: %d\n", packetType)
		}
	}
	// Clear the queue after processing
	nrs.queue = nrs.queue[:0]
}
