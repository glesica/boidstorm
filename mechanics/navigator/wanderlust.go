package navigator

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
)

func init() {
	Register(Wanderlust)
}

// Wanderlust pushes the individual away from or (more commonly, via a
// negative weight) toward its home location.
func Wanderlust(individual boid.T, neighbors []boid.T) vector.T {
	weight, location := individual.Config().Wanderlust()
	return individual.Position().To(location).Scale(weight)
}
