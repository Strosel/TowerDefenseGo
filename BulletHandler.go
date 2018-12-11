package main

import (
	"github.com/faiface/pixel/pixelgl"
)

// BulletHandler Handles drawing, moving and collisions of multiple Bullets
type BulletHandler []*Bullet

// Add Adds a Bullet to the set
func (bh *BulletHandler) Add(b *Bullet) {
	*bh = append(*bh, b)
}

// Merge Adds another BulletHandler into the set
func (bh *BulletHandler) Merge(bh2 BulletHandler) {
	*bh = append(*bh, bh2...)
}

// Move Moves the entire set of Bullets
func (bh BulletHandler) Move() {
	for _, b := range bh {
		b.Move()
	}
}

// Draw Draws the entire set of Bullets
func (bh BulletHandler) Draw(win *pixelgl.Window) {
	for _, b := range bh {
		b.Draw(win)
	}
}

// CheckCollision Checks if any Bullet has hit it's target, if so damages the target and destroys the bullet
func (bh *BulletHandler) CheckCollision() {
	for i, b := range *bh {
		if i < len(*bh) {
			if b.Alive() && b.Collision() {
				b.Target.Health -= b.Damage
				*bh = append((*bh)[:i], (*bh)[i+1:]...)
			} else if !b.Alive() {
				*bh = append((*bh)[:i], (*bh)[i+1:]...)
			}
		}
	}
}

// SelfDeath Kills the Bullet if it's Target is dead
func (bh *BulletHandler) SelfDeath() {
	for i, b := range *bh {
		if i < len(*bh) && !b.Target.Alive() {
			*bh = append((*bh)[:i], (*bh)[i+1:]...)
		}
	}
}
