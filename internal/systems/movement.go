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
	positionEntities := ms.world.GetEntitiesOfComponent(components.Position{})
	movementEntities := ms.world.GetEntitiesOfComponent(components.Move{})
	for e, move := range movementEntities {
		move := move.(*components.Move)
		direction := move.Direction
		position, ok := positionEntities[e]
		if !ok {
			continue
		}
		p := position.(*components.Position)
		switch direction {
		case components.None:
			continue
		case components.Down:
			p.Y += move.Speed
		case components.Up:
			p.Y -= move.Speed
		case components.Left:
			p.X -= move.Speed
		case components.Right:
			p.X += move.Speed
		default:
		}

		// reset movement
		if !move.Continuous {
			move.Direction = components.None
		}
	}
}
