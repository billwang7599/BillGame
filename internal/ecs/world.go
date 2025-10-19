package ecs

import "reflect"

// World is the central orchestrator, containing all the managers.
type World struct {
	entities   *EntityManager
	components *ComponentManager
	systems    *SystemManager
	Running    bool
}

// NewWorld creates an initialized World with all its managers ready to go.
func NewWorld() *World {
	return &World{
		entities:   NewEntityManager(),
		components: NewComponentManager(),
		systems:    NewSystemManager(),
		Running:    true,
	}
}

// --- Facade Methods ---
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
