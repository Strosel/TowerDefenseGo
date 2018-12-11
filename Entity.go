package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Drawable Defines a basic interface for Enteties
type Drawable interface {
	Draw(*pixelgl.Window)
	Bounds() pixel.Rect
}

// Entity A generic entity to be displayed
type Entity struct {
	Position, Scale, Size pixel.Vec
	Sprite                *pixel.Sprite
}

// NewEntity Creates a new Entity
func NewEntity(pos, size pixel.Vec, sprite pixel.Picture) *Entity {
	return &Entity{
		Position: pos,
		Scale:    pixel.V(size.X/sprite.Bounds().W(), size.Y/sprite.Bounds().H()),
		Size:     size,
		Sprite:   pixel.NewSprite(sprite, sprite.Bounds()),
	}
}

// Draw Draws the Entity's Sprite at Position
func (e Entity) Draw(win *pixelgl.Window) {
	mat := pixel.IM
	mat = mat.Moved(e.Position)
	mat = mat.ScaledXY(e.Position, e.Scale)
	e.Sprite.Draw(win, mat)
}

// Bounds Get the bounding rectangle of entity
func (e Entity) Bounds() pixel.Rect {
	TopRight := e.Position.Add(e.Size.Scaled(0.5))
	BottomLeft := e.Position.Sub(e.Size.Scaled(0.5))
	return pixel.R(BottomLeft.X, BottomLeft.Y, TopRight.X, TopRight.Y)
}

// DrawableRect Create a simple colored rectangle
func DrawableRect(rect pixel.Rect, color color.RGBA) *pixel.PictureData {
	pic := pixel.MakePictureData(rect)
	for i := range pic.Pix {
		pic.Pix[i] = color
	}
	return pic
}
