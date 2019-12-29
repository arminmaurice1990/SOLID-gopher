package entity

import "github.com/google/uuid"

type Entity = uuid.UUID

func NewEntity() Entity {
	return Entity(uuid.New())
}