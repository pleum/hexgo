package domain

import "time"

type Model interface {
	GetID() string
	SetID(id string)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
}

type Base struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Base) GetID() string {
	return b.ID
}

func (b *Base) SetID(id string) {
	b.ID = id
}

func (b *Base) SetCreatedAt(createdAt time.Time) {
	b.CreatedAt = createdAt
}

func (b *Base) SetUpdatedAt(updatedAt time.Time) {
	b.UpdatedAt = updatedAt
}

type Game struct {
	*Base
	Players []Player
}

type Player struct {
	*Base
	Name string
}
