package main

import (
	"math"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

// Tower Defines a tower with a Radius of aim and damage
type Tower struct {
	*Entity
	Radius    float64
	Damage    int
	Frequency int
}

// NewTower Create a new Tower
func NewTower(pos, size pixel.Vec, sprite pixel.Picture, r float64, dmg, frq int) *Tower {
	e := NewEntity(pos, size, sprite)
	return &Tower{
		Entity:    e,
		Radius:    r,
		Damage:    dmg,
		Frequency: frq,
	}
}

// Contains Checks if a point vec is inside the Tower's aimable radius
func (t Tower) Contains(vec pixel.Vec) bool {
	delta := t.Position.Sub(vec)
	return math.Sqrt(math.Pow(delta.X, 2)+math.Pow(delta.Y, 2)) <= t.Radius
}

// CanFire Checks if the Tower can fire in the current frame at the current frequency
func (t Tower) CanFire(frame int) bool {
	return frame%t.Frequency == 0
}

// FireAt Creates a bullet aimed at an Enemy
func (t Tower) FireAt(e *Enemy) *Bullet {
	pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Hotpink)
	return NewBullet(t.Position, pixel.V(50, 50), pic, e.Speed/10, e, t.Damage)
}
