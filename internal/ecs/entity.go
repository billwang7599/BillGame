package ecs

type Entity uint64

type EntityManager struct {
	nextID  Entity
	freeIDs []Entity
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextID:  1,
		freeIDs: make([]Entity, 0),
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

	return id
}

func (em *EntityManager) DeleteEntity(entity Entity, cm *ComponentManager) {
	em.freeIDs = append(em.freeIDs, entity)
	for _, entityComponentMapping := range cm.components {
		delete(entityComponentMapping, entity)
	}
}
