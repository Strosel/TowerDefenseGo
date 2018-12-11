package main

import "github.com/faiface/pixel"

// MovEntity A generic moving Entity
type MovEntity struct {
	*Entity
	Speed float64
}

// NewMovEntity Creates a new MovEntity
func NewMovEntity(pos, size pixel.Vec, sprite pixel.Picture, s float64) *MovEntity {
	e := NewEntity(pos, size, sprite)
	return &MovEntity{
		Entity: e,
		Speed:  s,
	}
}
