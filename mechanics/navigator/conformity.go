package navigator

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
)

func init() {
	Register(Conformity)
}

func Conformity(individual boid.T, neighbors []boid.T) vector.T {
	weight := individual.Config().Conformity()
	centroid := VelocityCentroid(neighbors)
	return centroid.Scale(weight)
}
