package main

import "github.com/faiface/pixel"

var (
	// Right Movement to the Right
	Right = pixel.V(1, 0)
	// Left Movement to the Left
	Left = pixel.V(-1, 0)
	// Up Movement Upward
	Up = pixel.V(0, 1)
	// Down Movement Downward
	Down = pixel.V(0, -1)
)

// Road Defines a Road, a block with a direction
type Road struct {
	*Entity
	Direction pixel.Vec
}

// NewRoad Create a new Road
func NewRoad(pos, size pixel.Vec, sprite pixel.Picture, dir pixel.Vec) *Road {
	e := NewEntity(pos, size, sprite)
	return &Road{
		Entity:    e,
		Direction: dir,
	}
}
