package mover

import "github.com/glesica/boidstorm/geometry/vector"

// T represents an object with position and velocity which allows us to
// run an mover against it to compute a new position.
type T interface {
	Position() vector.T
	Velocity() vector.T
}

// Next returns the next position for the given object and step size.
// TODO: Use the Verlet method for better stability
// TODO: Consider moving this to an "integrator" package
// TODO: Do we want this to be a method on an injectable struct?
// TODO: Should this return the vector to be added to the position instead? That way we could cap overall delta
// TODO: What about just returning an updated version with a Moved method? Would this work without generics?
// TODO: That would also suggest an Accelerated method the mechanics could use
func Next(ind T, step float64) vector.T {
	dv := ind.Velocity().Scale(step)
	return ind.Position().Add(dv)
}
