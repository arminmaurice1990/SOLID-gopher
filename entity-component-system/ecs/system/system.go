package system

import (
	"entity-component-system/ecs/entity"
	"entity-component-system/ecs/manager"
	"reflect"
)


type System interface {
	ListTypes() []reflect.Type
	GetSystemEntities() []entity.Entity
}

type BaseSystem struct {
	componentTypes []reflect.Type
	manager manager.Manager
}

func NewBaseSystem(componentTypes []reflect.Type, manager manager.Manager) *BaseSystem {
	return &BaseSystem{componentTypes:componentTypes, manager:manager}
}


func (bs *BaseSystem) ListTypes() []reflect.Type {
	return bs.componentTypes
}

func (bs BaseSystem) GetSystemEntities() []entity.Entity {
	return bs.manager.GetEntitiesWithComponents(bs.ListTypes())
}


