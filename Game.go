package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Game Defines a Game
type Game struct {
	Win                                    *pixelgl.Window
	Lvl                                    *Level
	Enemies                                *EnemyHandler
	Bullets                                *BulletHandler
	Buttons                                *ButtonHandler
	Frametxt, Wavetxt, Moneytxt, Healthtxt *TextEntity
	Frame, Wave, Money, Health             int
	Spawning, Debug                        bool
	StartTime                              time.Time
}

// NewGame Create a new Game
func NewGame(l *Level) *Game {
	return &Game{
		Lvl:     l,
		Enemies: &EnemyHandler{},
		Bullets: &BulletHandler{},
		Buttons: &ButtonHandler{},
		Frame:   0,
		Wave:    0,
		Money:   50, // get from levelsize?
		Health:  100,
	}
}

//Update Update the gamestate
func (g *Game) Update() {

	//Handle mouseclicks
	if g.Win.JustPressed(pixelgl.MouseButtonLeft) {
		g.MouseClick()
	}

	// Kill all Enemies that have left the Track and their Bullets, damage the player
	g.Health -= g.Enemies.SelfDeath(g.Lvl)
	g.Bullets.SelfDeath()

	// Move all the Enemies and Bullets
	g.Enemies.Move(g.Lvl)
	g.Bullets.Move()

	// Check if the Bullets kill their target, the player earns money
	g.Bullets.CheckCollision()
	g.Money += g.Enemies.CheckAlive()

	// Fire any new bullets
	newBullets := g.Lvl.Towers.Fire(g.Frame, g.Enemies)
	g.Bullets.Merge(newBullets)

	// Spawn any new enemies
	g.Spawning = g.Enemies.Spawn(g.Wave, g.Frame)

	g.Frame++
}

// MouseClick Handle mouse clicks
func (g *Game) MouseClick() {
	mouse := g.Win.MousePosition()

	// Is the player placing a tower?
	if PlaceTower != 0 && g.Lvl.GridIsEmpty(mouse) && g.Lvl.Bounds().Contains(mouse) {
		pos := g.Lvl.GridPos(mouse)
		pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Darkviolet)
		g.Lvl.AddTower(pos, pic, 250, 1, 10)
		g.Money -= PlaceTower
		PlaceTower = 0
	}

	// Did the player click a button?
	g.Buttons.Click(mouse)
}

//Draw Draw the current gamestate
func (g *Game) Draw() {
	if g.Health <= 0 {
		NewTextEntity(pixel.V(10, 350), 10, colornames.Red).Write("You Lost!", g.Win)
	} else if g.Wave > 20 {
		NewTextEntity(pixel.V(10, 350), 10, colornames.Green).Write("You Win!", g.Win)
	} else if g.Wave <= 20 {
		g.Lvl.Draw(g.Win)
		drawTmpTower(g.Win)
		g.Enemies.Draw(g.Win)
		g.Bullets.Draw(g.Win)
		g.Buttons.Draw(g.Win)

		if g.Debug {
			dt := time.Since(g.StartTime).Seconds()
			if dt < 1 {
				dt = 1
			}
			g.Frametxt.Write(fmt.Sprintf("avg. %vfps, frame: %v", g.Frame/int(dt), g.Frame), g.Win)
		}

		g.Wavetxt.Write(fmt.Sprintf("Wave: %v", g.Wave), g.Win)
		g.Moneytxt.Write(fmt.Sprintf("$%v", g.Money), g.Win)
		g.Healthtxt.Write(fmt.Sprintf("<3 %v", g.Health), g.Win)
	}
}
