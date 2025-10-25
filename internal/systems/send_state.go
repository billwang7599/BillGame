package systems

import (
	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
)

type SendStateSystem struct {
	world *ecs.World
}

func NewSendStateSystem(w *ecs.World) *SendStateSystem {
	return &SendStateSystem{
		world: w,
	}
}

func (sss *SendStateSystem) Update() {
	entityComponentMapping := sss.world.GetEntitiesOfComponent(components.Position{})
	for entity, component := range entityComponentMapping {
		position := component.(*components.Position)
		packet := ecs.EntityPositionPacket{
			Type:     ecs.EntityPositionPacketType,
			EntityId: entity,
			State:    *position,
		}
		sss.world.AddToPacketQueue(packet)
	}
}
