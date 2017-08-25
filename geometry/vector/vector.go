package vector

import "math"

type T interface {
	Add(other T) T
	Dist(other T) float64
	Magnitude() float64
	Rotate(angle float64) T
	Scale(scale float64) T
	To(other T) T
	X() float64
	Y() float64
}

// T is a 2-D vector type.
type t struct {
	x, y float64
}

// New returns a new vector with the given X and Y components.
func New(x, y float64) T {
	return &t{
		x: x,
		y: y,
	}
}

// Zero returns a new zero vector.
func Zero() T {
	return &t{
		x: 0.0,
		y: 0.0,
	}
}

// Add returns the vector that results from adding the target vector to
// the given vector.
func (v *t) Add(other T) T {
	return &t{
		x: v.x + other.X(),
		y: v.y + other.Y(),
	}
}

// Dist returns the distance from the target vector to the given vector.
func (v *t) Dist(other T) float64 {
	return v.To(other).Magnitude()
}

// Magnitude returns the magnitude of the target vector.
func (v *t) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

// Rotate returns the vector rotated by the given radians about the origin.
func (v *t) Rotate(angle float64) T {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return &t{
		x: v.x*cos - v.y*sin,
		y: v.x*sin + v.y*cos,
	}
}

// Scale returns the target vector multiplied by the given scalar.
func (v *t) Scale(scale float64) T {
	return &t{
		x: v.x * scale,
		y: v.y * scale,
	}
}

// To returns the vector from the target vector to the given vector.
func (v *t) To(other T) T {
	return &t{
		x: other.X() - v.x,
		y: other.Y() - v.y,
	}
}

// X returns the X component of the target vector.
func (v *t) X() float64 {
	return v.x
}

// Y returns the Y component of the target vector.
func (v *t) Y() float64 {
	return v.y
}
