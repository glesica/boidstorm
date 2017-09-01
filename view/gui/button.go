package gui

import "github.com/glesica/boidstorm/geometry/rect"

type Button interface {
	Widget

	// OnClick sets the click handler for this button.
	OnClick(func())
}

type button struct {
	bounds rect.T
}

func (b *button) Bounds() rect.T {
	
}
