package imdraw

import (
	"github.com/glesica/boidstorm/geometry/rect"
	"github.com/glesica/boidstorm/geometry/circle"
	"github.com/glesica/boidstorm/view"
	"github.com/glesica/boidstorm/geometry/line"
)

// TODO: Figure out how much bounds checking we want to do
// TODO: Implement bounds checking

type region struct {
	bounds rect.T
	canvas *imDrawCanvas
}

func (r *region) Active() bool {
	return r.canvas.Active()
}

func (r *region) Circle(shape circle.T, opts view.DrawOpts) {
	r.canvas.Circle(shape, opts)
}

func (r *region) Line(shape line.T, opts view.DrawOpts) {
	r.canvas.Line(shape, opts)
}

func (r *region) Rect(shape rect.T, opts view.DrawOpts) {
	r.canvas.Rect(shape, opts)
}

func (r *region) Region(bounds rect.T) view.Canvas {
	newBounds := rect.New(
		r.bounds.LowerLeft().X() + bounds.LowerLeft().X(),
		r.bounds.LowerLeft().Y() + bounds.LowerLeft().Y(),
		r.bounds.UpperRight().X() + bounds.UpperRight().X(),
			r.bounds.UpperRight().Y() + bounds.UpperRight().Y(),
	)
	return &region{
		bounds: newBounds,
		canvas: r.canvas,
	}
}

func (r *region) Size() rect.T {
	return r.bounds
}

func (r *region) Update() {
	r.canvas.Update()
}

