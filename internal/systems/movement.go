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
		speed := int64(move.Speed)
		switch direction {
		case components.None:
			continue
		case components.Down:
			p.Y += speed
		case components.Up:
			p.Y -= speed
		case components.Left:
			p.X -= speed
		case components.Right:
			p.X += speed
		default:
		}

		// reset movement
		if !move.Continuous {
			move.Direction = components.None
		}
	}
}
