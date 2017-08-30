package navigator

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/vector"
)

// TODO: Ideally we would weight the centroid calculation by distance from b

// PositionCentroid returns the centroid of the positions of the given
// slice of boids.
func PositionCentroid(boids []boid.T) vector.T {
	positions := make([]vector.T, len(boids))
	for i, neighbor := range boids {
		positions[i] = neighbor.Position()
	}

	return vector.Centroid(positions...)
}

// VelocityCentroid returns the centroid of the velocities of the given
// slice of boids.
func VelocityCentroid(boids []boid.T) vector.T {
	velocities := make([]vector.T, len(boids))
	for i, neighbor := range boids {
		velocities[i] = neighbor.Velocity()
	}

	return vector.Centroid(velocities...)
}
