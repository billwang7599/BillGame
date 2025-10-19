package systems

import (
	"github.com/billwang7599/BillGame/internal/components"
	"github.com/billwang7599/BillGame/internal/ecs"
	"github.com/gdamore/tcell/v2"
)

type RenderSystem struct {
	world         *ecs.World
	screen        tcell.Screen
	width, height int
}

func NewRenderSystem(world *ecs.World, screen tcell.Screen) *RenderSystem {
	width, height := screen.Size()
	return &RenderSystem{
		world:  world,
		screen: screen,
		width:  width,
		height: height,
	}
}

func (rs *RenderSystem) Update() {
	rs.screen.Clear()
	spriteMapping := rs.world.GetEntitiesOfComponent(components.Sprite{})
	positionMapping := rs.world.GetEntitiesOfComponent(components.Position{})
	for entity, sComponent := range spriteMapping {
		pComponent, ok := positionMapping[entity]
		if !ok {
			continue
		}
		sprite := sComponent.(*components.Sprite)
		position := pComponent.(*components.Position)
		rs.screen.SetContent(position.X, position.Y, sprite.Char, nil, tcell.StyleDefault)
	}
	rs.screen.Show()
}
