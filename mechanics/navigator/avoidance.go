package navigator

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
)

func init() {
	Register(Avoidance)
}

// Avoidance pushes the individual further from its neighbors.
func Avoidance(individual boid.T, neighbors []boid.T) vector.T {
	weight := individual.Config().Avoidance()
	centroid := PositionCentroid(neighbors)
	return individual.Position().To(centroid).Scale(-weight)
}
