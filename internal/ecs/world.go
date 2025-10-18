package ecs

// World is the central orchestrator, containing all the managers.
type World struct {
	entities   *EntityManager
	components *ComponentManager
	systems    *SystemManager
}

// NewWorld creates an initialized World with all its managers ready to go.
func NewWorld() *World {
	return &World{
		entities:   NewEntityManager(),
		components: NewComponentManager(),
		systems:    NewSystemManager(),
	}
}

// --- Facade Methods ---
// These methods provide a clean API and hide the internal managers.

func (w *World) NewEntity() Entity {
	return w.entities.NewEntity()
}

func (w *World) AddComponent(entity Entity, component any) {
	w.components.AddComponent(entity, component)
}

func (w *World) AddSystem(system System) {
	w.systems.Add(system)
}

func (w *World) GetComponents(component any) map[Entity]any {
	liveEntities := w.entities.GetLiveEntities()
	components := make(map[Entity]any)
	for _, e := range liveEntities {
		c, _ := w.components.GetComponent(e, component)
		if c != nil {
			components[e] = c
		}
	}
	return components
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
