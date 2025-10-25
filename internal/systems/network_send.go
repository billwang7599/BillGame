package systems

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"

	"github.com/billwang7599/BillGame/internal/ecs"
)

type NetworkSendSystem struct {
	conn  *net.UDPConn
	world *ecs.World
}

func NewSendSystem(world *ecs.World, serverAddrStr string) *NetworkSendSystem {
	serverAddr, err := net.ResolveUDPAddr("udp", serverAddrStr)
	if err != nil {
		log.Fatal(err)
	}
	world.Clients[serverAddr.String()] = 0

	system := NetworkSendSystem{
		conn:  world.Conn,
		world: world,
	}
	return &system
}

func (nss *NetworkSendSystem) Update() {
	if len(nss.world.PacketQueue) == 0 {
		return
	}

	for _, packet := range nss.world.PacketQueue {
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, packet)
		if err != nil {
			log.Printf("Failed to serialize packet: %v\n", err)
			continue
		}
		bytesToSend := buf.Bytes()
		for clientAddrStr := range nss.world.Clients {
			addr, err := net.ResolveUDPAddr("udp", clientAddrStr)
			if err != nil {
				log.Printf("Failed to resolve client address %v: %v", clientAddrStr, err)
				continue
			}
			_, err = nss.conn.WriteToUDP(bytesToSend, addr)
			if err != nil {
				log.Printf("Failed to send to %v: %v", clientAddrStr, err)
			} else {
				log.Printf("Sent packet of %d bytes to %v", len(bytesToSend), clientAddrStr)
			}
		}
	}

	// Clear the queue for the next tick
	nss.world.PacketQueue = nss.world.PacketQueue[:0]
}
