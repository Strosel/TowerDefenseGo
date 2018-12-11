package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

// TextEntity Generic text output entity
type TextEntity struct {
	Text *text.Text
	Size float64
}

// NewTextEntity Creates a new TextEntity
func NewTextEntity(pos pixel.Vec, size float64, clr color.Color) *TextEntity {
	t := text.New(pos, text.NewAtlas(basicfont.Face7x13, text.ASCII))
	t.Color = clr
	return &TextEntity{
		Text: t,
		Size: size,
	}
}

// Write Wite text to the entity
func (te TextEntity) Write(text string, win *pixelgl.Window) {
	te.Text.Clear()
	fmt.Fprintln(te.Text, text)
	te.Text.Draw(win, pixel.IM.Scaled(te.Text.Orig, te.Size))
}
