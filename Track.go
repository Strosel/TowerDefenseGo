package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Track Defines a Track, a collection of Roads and a start
type Track struct {
	Start *Road
	Path  map[pixel.Vec]*Road
}

// NewTrack Creates a new road from a string matrix
func NewTrack(h, w int, Map [][]string, ez pixel.Vec) Track {
	track := Track{}

	track.Path = map[pixel.Vec]*Road{}

	for _x, x := range Map {
		for _y, y := range x {
			if y == "R" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Brown)
				track.Start = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Right)
				track.Path[pixel.V(float64(_x), float64(_y))] = track.Start
			} else if y == "L" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Brown)
				track.Start = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Left)
				track.Path[pixel.V(float64(_x), float64(_y))] = track.Start
			} else if y == "U" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Brown)
				track.Start = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Up)
				track.Path[pixel.V(float64(_x), float64(_y))] = track.Start
			} else if y == "D" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Brown)
				track.Start = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Down)
				track.Path[pixel.V(float64(_x), float64(_y))] = track.Start
			} else if y == ">" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Green)
				track.Path[pixel.V(float64(_x), float64(_y))] = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Right)
			} else if y == "<" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Yellow)
				track.Path[pixel.V(float64(_x), float64(_y))] = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Left)
			} else if y == "^" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Pink)
				track.Path[pixel.V(float64(_x), float64(_y))] = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Up)
			} else if y == "V" {
				pic := DrawableRect(pixel.R(0, 0, 10, 10), colornames.Orange)
				track.Path[pixel.V(float64(_x), float64(_y))] = NewRoad(pixel.V(float64(_x), float64(_y)).ScaledXY(ez).Add(ez.Scaled(0.5)), ez, pic, Down)
			}
		}
	}

	return track
}

// Draw Draw the entire set of Roads
func (t Track) Draw(win *pixelgl.Window) {
	for _, r := range t.Path {
		r.Draw(win)
	}
}
