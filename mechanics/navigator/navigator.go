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

// Adjustments returns the adjustments to be applied to all boids during
// their navigation phase.
// TODO: Want to make sure the slice isn't mutable, return a copy
func Adjustments() []Adjustment {
	return adjustments
}

// Register causes the given adjustment to be applied to all boids during
// their navigation phase.
func Register(adjustment Adjustment) {
	adjustments = append(adjustments, adjustment)
}
