package vector

import "math"

// T is a 2-D vector type.
type T struct {
	x, y float64
}

// New returns a new vector with the given X and Y components.
func New(x, y float64) *T {
	return &T{
		x: x,
		y: y,
	}
}

// Zero returns a new zero vector.
func Zero() *T {
	return &T{
		x: 0.0,
		y: 0.0,
	}
}

// Add returns the vector that results from adding the target vector to
// the given vector.
func (v *T) Add(other *T) *T {
	return &T{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

// Dist returns the distance from the target vector to the given vector.
func (v *T) Dist(other *T) float64 {
	return v.To(other).Magnitude()
}

// Magnitude returns the magnitude of the target vector.
func (v *T) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

// Rotate returns the vector rotated by the given radians.
func (v *T) Rotate(angle float64) *T {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return &T{
		x: v.x*cos - v.y*sin,
		y: v.x*sin + v.y*cos,
	}
}

// Scale returns the target vector multiplied by the given scalar.
func (v *T) Scale(scale float64) *T {
	return &T{
		x: v.x * scale,
		y: v.y * scale,
	}
}

// To returns the vector from the target vector to the given vector.
func (v *T) To(other *T) *T {
	return &T{
		x: other.x - v.x,
		y: other.y - v.y,
	}
}

// X returns the X component of the target vector.
func (v *T) X() float64 {
	return v.x
}

// Y returns the Y component of the target vector.
func (v *T) Y() float64 {
	return v.y
}
