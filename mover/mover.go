package mover

import "github.com/glesica/boidstorm/geometry/vector"

type Integrator func(object T, step float64) vector.T

// T represents an object with position and velocity which allows us to
// run an mover against it to compute a new position.
type T interface {
	//Acceleration() vector.T
	Position() vector.T
	Velocity() vector.T
}

// Next returns the next position for the given object and step size.
// TODO: Use the Verlet method for better stability
func Next(individual T, step float64) vector.T {
	dv := individual.Velocity().Scale(step)
	return individual.Position().Add(dv)
}
