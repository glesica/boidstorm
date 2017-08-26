package boid

import (
	"github.com/glesica/boidstorm/geometry/circle"
	"github.com/glesica/boidstorm/geometry/line"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/view"
	"golang.org/x/image/colornames"
)

// We should attempt to do this in a way that enforces immutability, but we
// may want to cheat for performance reasons later, in which case we can just
// mutate and return.

// TODO: Where do boid movement properties live?
// We could just store a pointer to a configuration object since we'll
// need to build user-generated configurations anyway. Where would that
// struct live? Probably in the boid package?

type T interface {
	Accelerate(velocity vector.T) T
	Config() Config
	Draw(frame view.Frame)
	Move(position vector.T) T
	Position() vector.T
	Rotate(angle float64) T
	Velocity() vector.T
}

type t struct {
	config   Config
	position vector.T
	velocity vector.T
}

func New(x, y float64, c Config) T {
	return &t{config: c, position: vector.New(x, y), velocity: vector.Zero()}
}

func (b *t) Accelerate(velocity vector.T) T {
	return &t{
		position: b.position,
		velocity: b.velocity.Add(velocity),
	}
}

func (b *t) Config() Config {
	return b.config
}

func (b *t) Draw(frame view.Frame) {
	c := circle.New(b.position.X(), b.position.Y(), 10)
	l := line.New(b.position, b.position.Add(b.velocity.Scale(100)))
	o := view.DrawOpts{StrokeColor: colornames.Green, StrokeWidth: 1}
	frame.Line(l, o)
	frame.Circle(c, o)
}

func (b *t) Move(position vector.T) T {
	return &t{
		position: position,
		velocity: b.velocity,
	}
}

func (b *t) Position() vector.T {
	return b.position
}

func (b *t) Rotate(angle float64) T {
	return &t{
		position: b.position,
		velocity: b.velocity.Rotate(angle),
	}
}

func (b *t) Velocity() vector.T {
	return b.velocity
}
