package navigation

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/swarm"
)

func UpdateBoid(s swarm.T, b boid.T) boid.T {
	neighbors := s.Neighbors(b, 1000)
	if len(neighbors) == 0 {
		return b
	}

	positions := make([]vector.T, len(neighbors))
	for i, neighbor := range neighbors {
		positions[i] = neighbor.Position()
	}
	centroid := vector.Centroid(positions...)
	b.Accelerate(b.Position().To(centroid).Scale(0.1))
	return b.Accelerate(vector.New(0, 1).Scale(0.01))
}
