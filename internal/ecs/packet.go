package ecs

import (
	"github.com/billwang7599/BillGame/internal/components"
)

const (
	PlayerInputPacketType uint8 = iota + 1
	EntityPositionPacketType
)

type PlayerInputPacket struct {
	Type     uint8
	Key      rune
	EntityId Entity
}

type EntityPositionPacket struct {
	Type     uint8
	EntityId Entity
	State    components.Position
}
