package ecs

import (
	"fmt"
	"reflect"
)

type ComponentManager struct {
	components map[reflect.Type]map[Entity]any
}

func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		components: make(map[reflect.Type]map[Entity]any),
	}
}

func (cm *ComponentManager) AddComponent(e Entity, component any) {
	componentType := reflect.TypeOf(component)
	if _, ok := cm.components[componentType.Elem()]; !ok {
		cm.components[componentType.Elem()] = make(map[Entity]any)
	}
	cm.components[componentType.Elem()][e] = component
}

func (cm *ComponentManager) GetComponent(component any) map[Entity]any {
	entityComponentMapping, ok := cm.components[reflect.TypeOf(component)]
	if !ok {
		return nil
	}

	return entityComponentMapping
}

func (cm *ComponentManager) RemoveComponentFromEntity(e Entity, component any) {
	compMapping, ok := cm.components[reflect.TypeOf(component)]
	if !ok {
		return
	}

	delete(compMapping, e)
}

func (cm *ComponentManager) PrintComponent() {
	fmt.Println(cm.components)
}
