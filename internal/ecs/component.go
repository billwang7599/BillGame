package ecs

import (
	"errors"
	"reflect"
)

type ComponentManager struct {
	components map[Entity]map[reflect.Type]any
}

func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		components: make(map[Entity]map[reflect.Type]any),
	}
}

func (cm *ComponentManager) AddComponent(e Entity, comp any) {
	compType := reflect.TypeOf(comp)
	if _, ok := cm.components[e]; !ok {
		cm.components[e] = make(map[reflect.Type]any)
	}
	cm.components[e][compType] = comp
}

func (cm *ComponentManager) GetComponent(e Entity, comp any) (any, error) {
	compMapping, ok := cm.components[e]
	if !ok {
		return nil, errors.New("Entity doesn't exist")
	}

	component, ok := compMapping[reflect.TypeOf(comp)]
	if !ok {
		return nil, errors.New("Component doesn't exist")
	}

	return component, nil
}

func (cm *ComponentManager) RemoveComponent(e Entity, comp any) {
	compMapping, ok := cm.components[e]
	if !ok {
		return
	}
	delete(compMapping, reflect.TypeOf(comp))

}
