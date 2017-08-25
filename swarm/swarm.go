package swarm

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/rect"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/view"
	"math/rand"
)

// TODO: Should we make this immutable?
type T struct {
	boids []*boid.T
}

func New() *T {
	return &T{boids: make([]*boid.T, 0)}
}

func Random(size int, region rect.T) *T {
	xOffset := region.LowerLeft().X()
	xDelta := region.UpperRight().X() - region.LowerLeft().X()
	yOffset := region.LowerLeft().Y()
	yDelta := region.UpperRight().Y() - region.LowerLeft().Y()
	s := New()
	for i := 0; i < size; i++ {
		x := rand.Float64()*xDelta + xOffset
		y := rand.Float64()*yDelta + yOffset
		b := boid.New(x, y)
		a := vector.New(rand.Float64()*2-1.0, rand.Float64()*2-1.0)
		s.Add(b.Accelerated(a))
	}
	return s
}

// Add inserts the given boid into the swarm. If the boid already exists in
// the swarm it will not be added again.
// TODO: We should store a proper set to speed up the membership check
func (s *T) Add(ind *boid.T) {
	for _, b := range s.boids {
		if b == ind {
			return
		}
	}
	s.boids = append(s.boids, ind)
}

func (s *T) ForEach(callback func(b *boid.T)) {
	for _, b := range s.boids {
		callback(b)
	}
}

func (s *T) Draw(frame view.Frame) {
	for _, b := range s.boids {
		b.Draw(frame)
	}
}

// Near returns a slice of the boids from this swarm that are within the
// given distance of the given position.
// TODO: We should store something like a k-D tree to make this faster
func (s *T) Near(position vector.T, distance float64) []*boid.T {
	near := make([]*boid.T, 0)
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
func (s *T) Neighbors(ind *boid.T, distance float64) []*boid.T {
	near := s.Near(ind.Position(), distance)
	for i, b := range near {
		if b == ind {
			near = append(near[:i], near[i+1:]...)
		}
	}
	return near
}

func (s *T) Update(callback func(b *boid.T) *boid.T) {
	for i, b := range s.boids {
		s.boids[i] = callback(b)
	}
}
