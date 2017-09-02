package swarm

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/rect"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/view"
	"math/rand"
)

type T interface {
	Add(newbie boid.T) T
	ForEach(callback func(b boid.T))
	Draw(canvas view.Canvas)
	Near(position vector.T, distance float64) []boid.T
	Neighbors(individual boid.T, distance float64) []boid.T
	Update(callback func(b boid.T) boid.T) T
}

type t struct {
	boids []boid.T
}

func New() T {
	return &t{boids: make([]boid.T, 0)}
}

func Random(size int, region rect.T, config boid.Config) T {
	xOffset := region.LowerLeft().X()
	xDelta := region.UpperRight().X() - region.LowerLeft().X()
	yOffset := region.LowerLeft().Y()
	yDelta := region.UpperRight().Y() - region.LowerLeft().Y()
	s := New()
	for i := 0; i < size; i++ {
		x := rand.Float64()*xDelta + xOffset
		y := rand.Float64()*yDelta + yOffset
		b := boid.New(x, y, config)
		a := vector.New(
			rand.Float64()*2*config.MaxSpeed()-config.MaxSpeed(),
			rand.Float64()*2*config.MaxSpeed()-config.MaxSpeed(),
		)
		s = s.Add(b.Accelerate(a))
	}
	return s
}

// Add inserts the given boid into the swarm. If the boid already exists in
// the swarm it will not be added again.
// TODO: We should store a proper set to speed up the membership check
func (s *t) Add(newbie boid.T) T {
	for _, b := range s.boids {
		if b == newbie {
			return s
		}
	}

	newBoids := make([]boid.T, len(s.boids)+1)
	copy(newBoids, s.boids)
	newBoids[len(newBoids)-1] = newbie

	return &t{
		boids: newBoids,
	}
}

func (s *t) ForEach(callback func(b boid.T)) {
	for _, b := range s.boids {
		callback(b)
	}
}

func (s *t) Draw(canvas view.Canvas) {
	for _, b := range s.boids {
		b.Draw(canvas)
	}
}

// Near returns a slice of the boids from this swarm that are within the
// given distance of the given position.
// TODO: We should store something like a k-D tree to make this faster
func (s *t) Near(position vector.T, distance float64) []boid.T {
	near := make([]boid.T, 0)
	for _, b := range s.boids {
		if position.Dist(b.Position()) <= distance {
			near = append(near, b)
		}
	}
	return near
}

// Neighbors returns a slice of the boids that are within the given distance
// from the given boid. The boid provided needn't actually be a member of the
// swarm, but if it is it will be removed from the resulting slice.
func (s *t) Neighbors(individual boid.T, distance float64) []boid.T {
	near := s.Near(individual.Position(), distance)
	for i, b := range near {
		if b == individual {
			near = append(near[:i], near[i+1:]...)
		}
	}
	return near
}

func (s *t) Update(callback func(b boid.T) boid.T) T {
	newBoids := make([]boid.T, len(s.boids))

	for i, b := range s.boids {
		newBoids[i] = callback(b)
	}

	return &t{
		boids: newBoids,
	}
}
