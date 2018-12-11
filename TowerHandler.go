package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// TowerHandler Handles drawing and fireing of multiple towers
type TowerHandler map[pixel.Vec]*Tower

// Draw Draw the entire set of Towers
func (th TowerHandler) Draw(win *pixelgl.Window) {
	for _, t := range th {
		t.Draw(win)
	}
}

// Fire Checks all Towers and Enemies to fire Bullets at Enemies within range
func (th TowerHandler) Fire(frame int, eh *EnemyHandler) BulletHandler {
	var bh BulletHandler

	for _, t := range th {
		if t.CanFire(frame) {
			for _, e := range eh.Enemies {
				if t.Contains(e.Position) {
					b := t.FireAt(e)
					bh = append(bh, b)
					break
				}
			}
		}
	}

	return bh
}
