package main

import (
	"image"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	game *Game

	lvl                                    Level
	enemies                                EnemyHandler
	bullets                                BulletHandler
	buttons                                ButtonHandler
	frametxt, wavetxt, moneytxt, healthtxt *TextEntity
	t                                      time.Time

	frame  = 0
	wave   = 0
	money  = 0
	health = 100
	// PlaceTower Describes if the player is currently placing a Tower
	PlaceTower   = 0
	windowWidth  = 700.0
	windowHeight = 700.0
	levelWidth   = 500
	levelHeight  = 500
	sprites      = map[string]pixel.Picture{}
	spawning     = false
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func drawTmpTower(win *pixelgl.Window) {
	if PlaceTower != 0 {
		mouse := win.MousePosition()
		tcolor := colornames.Darkviolet
		if !game.Lvl.GridIsEmpty(mouse) || !game.Lvl.Bounds().Contains(mouse) {
			tcolor = colornames.Maroon
		}
		tcolor.A /= 2
		pic := DrawableRect(pixel.R(0, 0, 10, 10), tcolor)
		NewTower(mouse, game.Lvl.EntitySize, pic, 0, 0, 0).Draw(win)
	}
}

func run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Tower Defense",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	game.Win = win

	for !win.Closed() {
		win.Clear(colornames.Floralwhite)

		game.Update()
		//update(win)
		game.Draw()
		//draw(win)

		win.Update()
	}
}

func main() {
	sprites["tower"] = DrawableRect(pixel.R(0, 0, 10, 10), colornames.Darkviolet)
	sprites["enemy"] = DrawableRect(pixel.R(0, 0, 10, 10), colornames.Cyan)

	game = NewGame(NewLevel(windowHeight, levelHeight, levelWidth, "./level.txt"))

	game.Enemies.OnSpawn = func(eh *EnemyHandler, wave int, frame int) bool {
		if eh.LastWave < wave && eh.WaveStart == 0 {
			eh.WaveStart = frame
			eh.LastWave = wave
		}
		if eh.LastWave == wave {
			if wave == 1 && frame-eh.WaveStart < 100 { // if running wave one and not past duration of 100 frames
				if frame-eh.WaveStart == 0 {
					eh.Add(NewEnemy(game.Lvl.Track.Start.Position, game.Lvl.EntitySize, sprites["enemy"], 1, 1))
				} else if frame-eh.WaveStart == 60 {
					eh.Add(NewEnemy(game.Lvl.Track.Start.Position, game.Lvl.EntitySize, sprites["enemy"], 1, 10))
				}
				return true // we are running a wave
			}
		}
		return false
	}

	pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Gray)
	game.Buttons.Add(NewButton(pixel.V(30, 30), pixel.V(50, 50), pic, func() {
		cost := 10
		if PlaceTower == cost { // untoggle purchase
			PlaceTower = 0
		} else if game.Money >= cost { // able to purchase
			PlaceTower = cost
		}
	}))

	pic = DrawableRect(pixel.R(0, 0, 10, 10), colornames.Green)
	game.Buttons.Add(NewButton(pixel.V(85, 30), pixel.V(50, 50), pic, func() {
		if !game.Spawning {
			game.Wave++
		}
	}))

	game.Frametxt = NewTextEntity(pixel.V(10, 20), 1, colornames.Black)
	game.Wavetxt = NewTextEntity(pixel.V(10, 10), 1, colornames.Black)
	game.Moneytxt = NewTextEntity(pixel.V(510, 620), 2, colornames.Green)
	game.Healthtxt = NewTextEntity(pixel.V(510, 670), 2, colornames.Red)

	game.StartTime = time.Now()
	game.Debug = true

	pixelgl.Run(run)
}
