package main

import (
	"github.com/faiface/pixel"
)

// Bullet Defines a bullet that targets and damages an Entity
type Bullet struct {
	*MovEntity
	Target *Enemy
	Damage int
}

// NewBullet Creates a new Bullet
func NewBullet(pos, size pixel.Vec, sprite pixel.Picture, s float64, en *Enemy, dmg int) *Bullet {
	e := NewMovEntity(pos, size, sprite, s)
	return &Bullet{
		MovEntity: e,
		Target:    en,
		Damage:    dmg,
	}
}

// Move moves the Bullet towards its target
func (b *Bullet) Move() {
	diff := b.Target.Position.Sub(b.Position)
	b.Position = b.Position.Add(diff.Scaled(b.Speed))
}

// Collision Checks if the Bullet has hit it's target
func (b Bullet) Collision() bool {
	return b.Target.Bounds().Contains(b.Position)
}

// Alive Checks if the Bullet's target is alive
func (b Bullet) Alive() bool {
	if b.Target == nil {
		return false
	}
	return b.Target.Alive()
}
