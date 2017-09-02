package gui

import (
	"github.com/glesica/boidstorm/view"
	"github.com/glesica/boidstorm/geometry/rect"
)

// Thoughts
// The widget will always take the full canvas area it is given, or at
// least as much of it as it wants to occupy. It will then be the
// responsibility of the layout to figure out how much of the available
// canvas to allocate to each widget. If the provided canvas is too
// small for the widget to be drawn properly it should do the best it
// can, possibly clipping itself, or even drawing an entirely different
// version of itself.

type Widget interface {
	view.Drawer

	MaxSize() rect.T
	MinSize() rect.T
}
