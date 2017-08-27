package navigation

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/swarm"
)

// TODO: Consider accelerating toward the centroid of the current and the group
// This might moderate the corrections a bit automatically
// TODO: Ideally we would weight the centroid calculation by distance from b
func Boid(s swarm.T, b boid.T) boid.T {
	neighbors := s.Neighbors(b, 1000)
	if len(neighbors) == 0 {
		return b
	}

	positions := make([]vector.T, len(neighbors))
	velocities := make([]vector.T, len(neighbors))
	for i, neighbor := range neighbors {
		positions[i] = neighbor.Position()
		velocities[i] = neighbor.Velocity()
	}

	positionCentroid := vector.Centroid(positions...)
	velocityCentroid := vector.Centroid(velocities...)

	acceleration := vector.Zero()

	// Avoidance
	avoidance := b.Config().Avoidance()
	acceleration = acceleration.Add(b.Position().To(positionCentroid).Scale(-1 * avoidance))

	// Conformity
	conformity := b.Config().Conformity()
	acceleration = acceleration.Add(velocityCentroid.Scale(conformity))

	return b.Accelerate(acceleration.Clamp(0.01))
}
