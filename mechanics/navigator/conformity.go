package navigator

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
)

func init() {
	Register(Conformity)
}

// Conformity aligns the velocity of the individual with those of its
// neighbors.
func Conformity(individual boid.T, neighbors []boid.T) vector.T {
	weight := individual.Config().Conformity()
	centroid := VelocityCentroid(neighbors)
	return centroid.Scale(weight)
}
