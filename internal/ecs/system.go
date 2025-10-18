package ecs

type System interface {
	Update()
}

type SystemManager struct {
	systems []System
}

func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: make([]System, 0),
	}
}

func (sm *SystemManager) Add(system System) {
	sm.systems = append(sm.systems, system)
}

func (sm *SystemManager) Update() {
	for _, system := range sm.systems {
		system.Update()
	}
}
