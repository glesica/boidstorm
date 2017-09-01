package gui

import (
	"github.com/glesica/boidstorm/view"
	"github.com/glesica/boidstorm/geometry/rect"
)

type Widget interface {
	// Bounds returns the widget's bounding box.
	// TODO: Use something more efficient for hit detection
	Bounds() rect.T

	// Draw allows the widget to be drawn on a frame.
	Draw(frame view.Canvas)
}
