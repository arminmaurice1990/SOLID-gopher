package manager

import (
	"entity-component-system/ecs/component"
	"entity-component-system/ecs/entity"
	"errors"
	"reflect"
)

type Manager interface {
	AddEntity(entity entity.Entity) error
	RemoveEntity(entity entity.Entity)
	GetEntitiesWithComponents(compTypes []reflect.Type) []entity.Entity
}


type manager struct {
	entityrelations map[entity.Entity][]component.Component
}

func NewManager() *manager {
	return &manager{entityrelations: map[entity.Entity][]component.Component{}}
}

func (m *manager) AddEntity(entity entity.Entity) error {
	for k, _ := range m.entityrelations {
		if k == entity {
			return errors.New("Entity already exists")
		}
	}
	m.entityrelations[entity] = []component.Component{}
	return nil
}

func (m *manager) RemoveEntity(entity entity.Entity) {
	delete(m.entityrelations, entity)
}

func (m *manager) GetEntitiesWithComponents (queryTypes []reflect.Type) []entity.Entity {
	entites := []entity.Entity{}
	for k, v := range m.entityrelations {
		checks := len(queryTypes)
		for _, compType := range v {
			for _, queryType := range queryTypes {
				if queryType == compType.GetType() {
					checks--
				}
			}
			if checks < 1 {
				entites = append(entites, k)
			}
		}
	}

	return entites
}