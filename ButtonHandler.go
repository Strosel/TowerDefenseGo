package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// ButtonHandler Handles the clicking of drawing of multiple Buttons
type ButtonHandler []*Button

// Add Adds a Button to the set
func (bh *ButtonHandler) Add(b *Button) {
	*bh = append(*bh, b)
}

// Draw Draw the entire set fo Buttons
func (bh ButtonHandler) Draw(win *pixelgl.Window) {
	for _, b := range bh {
		b.Draw(win)
	}
}

// Click Check if the given position aplies to any button and clicks it
func (bh ButtonHandler) Click(click pixel.Vec) {
	for _, b := range bh {
		b.OnClick(click)
	}
}
