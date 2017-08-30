package navigator

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
)

// An Adjustment is a function that accepts a boid and a set of neighboring
// boids and returns an acceleration vector to be applied (possibly after
// scaling) to the given individual boid.
type Adjustment func(individual boid.T, neighbors []boid.T) vector.T

var adjustments []Adjustment = make([]Adjustment, 0)

// Apply applies to registered adjustments to the given individual and
// neighbors.
// TODO: This is maybe where normalization should happen
func Apply(individual boid.T, neighbors []boid.T) vector.T {
	acceleration := vector.Zero()
	for _, adjustment := range adjustments {
		acceleration = acceleration.Add(adjustment(individual, neighbors))
	}
	return acceleration
}

// Register causes the given adjustment to be applied to all boids during
// their navigation phase.
func Register(adjustment Adjustment) {
	adjustments = append(adjustments, adjustment)
}
