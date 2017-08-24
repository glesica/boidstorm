package mover_test

import (
	"github.com/glesica/boidstorm/mover"
	"github.com/glesica/boidstorm/test"
	"github.com/glesica/boidstorm/vector"
	"testing"
)

type stubMover struct {
	position *vector.T
	velocity *vector.T
}

func (m *stubMover) Velocity() *vector.T {
	return m.velocity
}

func (m *stubMover) Position() *vector.T {
	return m.position
}

func TestNext(t *testing.T) {
	m0 := &stubMover{
		position: vector.Zero(),
		velocity: vector.New(1.0, 1.0),
	}
	p0 := mover.Next(m0, 0.5)
	e0 := vector.New(0.5, 0.5)
	test.ExpectEqualVectors(t, "Next(m0, 0.5)", p0, e0)
}
