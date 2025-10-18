package systems

import (
	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
)

type MovementSystem struct {
	world *ecs.World
}

func NewMovementSystem(world *ecs.World) *MovementSystem {
	return &MovementSystem{
		world: world,
	}
}

func (ms *MovementSystem) Update() {
	positionMapping := ms.world.GetComponents(components.NewPosition(0, 0))
	movementMapping := ms.world.GetComponents(components.NewMove(0))
	for e, move := range movementMapping {
		direction := move.Direction
	}
}
