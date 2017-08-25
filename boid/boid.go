package boid

import (
	"github.com/glesica/boidstorm/geometry/circle"
	"github.com/glesica/boidstorm/geometry/line"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/view"
	"golang.org/x/image/colornames"
)

type T struct {
	position vector.T
	velocity vector.T
}

func New(x, y float64) *T {
	return &T{position: vector.New(x, y), velocity: vector.Zero()}
}

func (b *T) Accelerated(vel vector.T) *T {
	return &T{
		position: b.position,
		velocity: vel,
	}
}

func (b *T) Draw(frame view.Frame) {
	c := circle.New(b.position.X(), b.position.Y(), 10)
	l := line.New(b.position, b.position.Add(b.velocity.Scale(100)))
	o := view.DrawOpts{StrokeColor: colornames.Green, StrokeWidth: 1}
	frame.Line(l, o)
	frame.Circle(c, o)
}

func (b *T) Moved(pos vector.T) *T {
	return &T{
		position: pos,
		velocity: b.velocity,
	}
}

func (b *T) Position() vector.T {
	return b.position
}

func (b *T) Velocity() vector.T {
	return b.velocity
}

// TODO: Methods to update position and velocity
// We should attempt to do this in a way that enforces immutability, but we
// may want to cheat for performance reasons later, in which case we can just
// mutate and return.
