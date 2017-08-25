package view

import (
	"github.com/glesica/boidstorm/geometry/circle"
	"github.com/glesica/boidstorm/geometry/line"
	"github.com/glesica/boidstorm/geometry/rect"
)

type Frame interface {
	Active() bool
	Circle(shape circle.T, opts DrawOpts)
	Line(shape line.T, opts DrawOpts)
	Rect(shape rect.T, opts DrawOpts)
	//Text(center *vector.T, text string, opts DrawOpts)
	Update()
}
