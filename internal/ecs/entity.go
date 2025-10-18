package ecs

type Entity uint64

type EntityManager struct {
	nextID       Entity
	freeIDs      []Entity
	liveEntities map[Entity]bool
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextID:       1,
		freeIDs:      make([]Entity, 0),
		liveEntities: make(map[Entity]bool),
	}
}

// NewEntity gets an ID from the free list or creates a new one.
func (em *EntityManager) NewEntity() Entity {
	var id Entity
	if len(em.freeIDs) > 0 {
		// Pop an ID from the free list
		lastIndex := len(em.freeIDs) - 1
		id = em.freeIDs[lastIndex]
		em.freeIDs = em.freeIDs[:lastIndex]
	} else {
		id = em.nextID
		em.nextID++
	}

	// If no IDs are in the free list, create a new one
	em.liveEntities[id] = true
	return id
}

func (em *EntityManager) DeleteEntity(entity Entity, cm *ComponentManager) {
	em.freeIDs = append(em.freeIDs, entity)
	delete(cm.components, entity)
	delete(em.liveEntities, entity)
}

func (em *EntityManager) GetLiveEntities() []Entity {
	sum := make([]Entity, 0, len(em.liveEntities))
	for key, val := range em.liveEntities {
		if val {
			sum = append(sum, key)
		}
	}
	return sum
}
