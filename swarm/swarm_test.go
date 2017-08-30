package swarm_test

import (
	"testing"
	"github.com/glesica/boidstorm/swarm"
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/rect"
)

func expectSwarmSize(t *testing.T, s swarm.T, expectedSize int) {
	size := 0
	s.ForEach(func(b boid.T) {
		size += 1
	})
	if size != expectedSize {
		t.Errorf("Expected swarm size %d, found %d", expectedSize, size)
	}
}

func TestNew(t *testing.T) {
	s := swarm.New()
	t.Run("swarm should start empty", func(t *testing.T) {
		expectSwarmSize(t, s, 0)
	})
}

func TestRandom(t *testing.T) {
	size := 1000
	s := swarm.Random(size, rect.New(0, 0, 100, 100), boid.NewConfig())
	t.Run("Random swarm should have correct size", func(t *testing.T) {
		expectSwarmSize(t, s, size)
	})

	// TODO: Test that all are in bounds
}

func TestT_Add(t *testing.T) {
	s := swarm.New()
	t.Run("Add method should add to swarm", func(t *testing.T) {
		s = s.Add(boid.New(0, 0, nil))
		expectSwarmSize(t, s, 1)
	})
}
