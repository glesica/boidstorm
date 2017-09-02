package boid

import (
	"github.com/glesica/boidstorm/geometry/circle"
	"github.com/glesica/boidstorm/geometry/line"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/mover"
	"github.com/glesica/boidstorm/view"
	"golang.org/x/image/colornames"
)

// We should attempt to do this in a way that enforces immutability, but we
// may want to cheat for performance reasons later, in which case we can just
// mutate and return.

type T interface {
	// Accelerate the boid along the given vector. The magnitude of the
	// boid's total velocity will be clamped to respect its max speed.
	Accelerate(acceleration vector.T) T

	// Config returns the configuration for the target boid.
	Config() Config

	// Draw the target boid onto the given canvas.
	Draw(canvas view.Canvas)

	// Move updates the target boid's position using the given integrator.
	Move(intergrator mover.Integrator, step float64) T

	// Position returns the target boid's current position.
	Position() vector.T

	// Velocity returns the target boid's current velocity.
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

func (b *t) Accelerate(acceleration vector.T) T {
	newVelocity := b.velocity.RotateToward(
		acceleration, 0.01).Add(
		acceleration.Scale(0.5)).Clamp(b.config.MaxSpeed())
	return &t{
		config:   b.config,
		position: b.position,
		//velocity: b.velocity.Add(acceleration).Clamp(b.config.MaxSpeed()),
		velocity: newVelocity,
	}
}

func (b *t) Config() Config {
	return b.config
}

func (b *t) Draw(canvas view.Canvas) {
	c := circle.New(b.position.X(), b.position.Y(), 10)
	l := line.New(b.position, b.position.Add(b.velocity.Scale(5)))
	o := view.DrawOpts{StrokeColor: colornames.Green, StrokeWidth: 1}
	canvas.Line(l, o)
	canvas.Circle(c, o)
}

func (b *t) Move(integrator mover.Integrator, step float64) T {
	newPosition := integrator(b, step)
	return &t{
		config:   b.config,
		position: newPosition,
		velocity: b.velocity,
	}
}

func (b *t) Position() vector.T {
	return b.position
}

func (b *t) Velocity() vector.T {
	return b.velocity
}
