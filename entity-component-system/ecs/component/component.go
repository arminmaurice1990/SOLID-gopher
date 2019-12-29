package component

import (
	"entity-component-system/ecs/entity"
	"reflect"
)

type Component interface {
	GetType() reflect.Type
	GetEntity() entity.Entity
}

type BaseComponent struct {
	entity entity.Entity
}

func (bc *BaseComponent) GetType() reflect.Type {
	return reflect.TypeOf(bc)
}

func (bc *BaseComponent) GetEntity() entity.Entity {
	return bc.entity
}
