package main

import (
	"github.com/faiface/pixel"
)

// Button Defines a button
type Button struct {
	*Entity
	onClickDo func()
}

// NewButton create a new button
func NewButton(pos, size pixel.Vec, sprite pixel.Picture, oc func()) *Button {
	e := NewEntity(pos, size, sprite)
	return &Button{
		Entity:    e,
		onClickDo: oc,
	}
}

// OnClick Run the OnClickDo function if eligeble
func (b Button) OnClick(click pixel.Vec) {
	if b.Bounds().Contains(click) {
		b.onClickDo()
	}
}
