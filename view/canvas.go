package view

import (
	"github.com/glesica/boidstorm/geometry/circle"
	"github.com/glesica/boidstorm/geometry/line"
	"github.com/glesica/boidstorm/geometry/rect"
)

type Canvas interface {
	Active() bool

	Circle(shape circle.T, opts DrawOpts)

	Line(shape line.T, opts DrawOpts)

	Rect(shape rect.T, opts DrawOpts)

	Region(bounds rect.T) Canvas

	// TODO: Need a way to add and control text
	//Text(center *vector.T, text string, opts DrawOpts)

	Size() rect.T

	Update()
}
