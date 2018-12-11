package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// EnemyHandler Handles drawing, moving and killing of multiple Enemies
type EnemyHandler struct {
	Enemies             []*Enemy
	OnSpawn             func(*EnemyHandler, int, int) bool
	WaveStart, LastWave int
}

// Add Adds an Enemy to the set
func (eh *EnemyHandler) Add(e *Enemy) {
	eh.Enemies = append(eh.Enemies, e)
}

// Move Moves the entire set of Enemies
func (eh EnemyHandler) Move(level *Level) {
	for _, e := range eh.Enemies {
		e.Move(level)
	}
}

// Draw Draws the entire set of Enemies
func (eh EnemyHandler) Draw(win *pixelgl.Window) {
	for _, e := range eh.Enemies {
		e.Draw(win)
	}
}

// SelfDeath Kill all Enemies outside of the map and returns the damage dealt to the player
func (eh *EnemyHandler) SelfDeath(l *Level) int {
	dmg := 0
	for _, e := range eh.Enemies {
		if !l.Bounds().Contains(e.Position) && (e.Bounds().Intersect(l.Track.Start.Bounds())) == pixel.R(0, 0, 0, 0) {
			dmg += e.Health
			e.Health = 0
		}
	}
	eh.CheckAlive()
	return dmg
}

// CheckAlive Destroys all dead Enemies, returns the oney earned by the player
func (eh *EnemyHandler) CheckAlive() int {
	money := 0
	for i, e := range eh.Enemies {
		if !e.Alive() {
			money += int(e.Speed)
			eh.Enemies = append(eh.Enemies[:i], eh.Enemies[i+1:]...)
		}
	}
	return money
}

// Spawn Span Enemies according to the wave
func (eh *EnemyHandler) Spawn(wave, frame int) bool {
	return eh.OnSpawn(eh, wave, frame)
}
