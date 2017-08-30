package navigator

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
)

func init() {
	Register(Wanderlust)
}

func Wanderlust(individual boid.T, neighbors []boid.T) vector.T {
	weight, location := individual.Config().Home()
	return individual.Position().To(location).Scale(-weight)
}
