package main

import (
	"fmt"

	"github.com/faiface/pixel"
)

// Enemy Defines an Enemy that moves via a track in the Level
type Enemy struct {
	*MovEntity
	Health  int
	LastDir pixel.Vec
}

// NewEnemy Creates a new Enemy
func NewEnemy(pos, size pixel.Vec, sprite pixel.Picture, s float64, hp int) *Enemy {
	e := NewMovEntity(pos, size, sprite, s)
	return &Enemy{
		MovEntity: e,
		Health:    hp,
	}
}

// Origin Gets the origin
func (e Enemy) Origin() pixel.Vec {
	return e.Position
}

// Top Gets the middle of the Top side
func (e Enemy) Top() pixel.Vec {
	return e.Position.Add(e.Size.ScaledXY(pixel.V(0, 0.5)))
}

// Left Gets the middle of the Left side
func (e Enemy) Left() pixel.Vec {
	return e.Position.Sub(e.Size.ScaledXY(pixel.V(0.5, 0)))
}

// Bottom Gets the middle of the Bottom side
func (e Enemy) Bottom() pixel.Vec {
	return e.Position.Sub(e.Size.ScaledXY(pixel.V(0, 0.5)))
}

// Right Gets the middle of the Right side
func (e Enemy) Right() pixel.Vec {
	return e.Position.Add(e.Size.ScaledXY(pixel.V(0.5, 0)))
}

// Move Moves the enemy according to the Track in the Level
func (e *Enemy) Move(Level *Level) {
	var (
		currentRoad *Road
		err         error
	)

	if e.LastDir == Right {
		currentRoad, err = Level.GetRoad(e.Left())
		if err != nil {
			currentRoad, err = Level.GetRoad(e.Right())
			if err != nil {
				fmt.Println("Can't Move")
			}
		}
	} else if e.LastDir == Left {
		currentRoad, err = Level.GetRoad(e.Right())
		if err != nil {
			currentRoad, err = Level.GetRoad(e.Left())
			if err != nil {
				fmt.Println("Can't Move")
			}
		}
	} else if e.LastDir == Up {
		currentRoad, err = Level.GetRoad(e.Bottom())
		if err != nil {
			currentRoad, err = Level.GetRoad(e.Top())
			if err != nil {
				fmt.Println("Can't Move")
			}
		}
	} else if e.LastDir == Down {
		currentRoad, err = Level.GetRoad(e.Top())
		if err != nil {
			currentRoad, err = Level.GetRoad(e.Bottom())
			if err != nil {
				fmt.Println("Can't Move")
			}
		}
	} else {
		currentRoad, err = Level.GetRoad(e.Origin())
		if err != nil {
			fmt.Println("Can't Move")
		}
	}

	e.Position = e.Position.Add(currentRoad.Direction.Scaled(e.Speed))

	if e.LastDir != currentRoad.Direction {
		e.LastDir = currentRoad.Direction
		e.Position = currentRoad.Position
	}
}

// Alive Check if the Enemy is alive
func (e Enemy) Alive() bool {
	return e.Health > 0
}
