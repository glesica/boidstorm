package vector

import "math"

type T interface {
	// Add returns the vector that results from adding the target vector to
	// the given vector.
	Add(other T) T

	// Angle returns the vector's heading as an angle in radians. This
	// method assumes the standard coordinate plane (the zero angle vector
	// points due right).
	Angle() float64

	// AngleTo returns the difference in angles between the target and the
	// given vector. This is essentially the value with the smallest
	// magnitude that, if passed to `Rotate` would make the vectors equal.
	AngleTo(other T) float64

	// Clamp adjusts the vector such that its magnitude will be no more than
	// the given value.
	Clamp(magnitude float64) T

	// Dist returns the distance from the target vector to the given vector.
	Dist(other T) float64

	// Magnitude returns the magnitude of the target vector.
	Magnitude() float64

	// Rotate returns the vector rotated by the given radians about the origin.
	Rotate(angle float64) T

	// Toward returns a vector rotated toward another vector by no more than
	// the given max number of radians.
	RotateToward(other T, maxDelta float64) T

	// Scale returns the target vector multiplied by the given scalar.
	Scale(scale float64) T

	// To returns the vector from the target vector to the given vector.
	To(other T) T

	// X returns the X component of the target vector.
	X() float64

	// Y returns the Y component of the target vector.
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

func (v *t) Add(other T) T {
	return &t{
		x: v.x + other.X(),
		y: v.y + other.Y(),
	}
}

func (v *t) Angle() float64 {
	return math.Atan2(v.y, v.x)
}

func (v *t) AngleTo(other T) float64 {
	a0 := v.Angle()
	a1 := other.Angle()

	// TODO: This could probably be improved
	if math.Abs(a1-a0) <= math.Pi {
		return math.Mod(a1-a0, math.Pi)
	} else {
		return math.Mod(a0-a1, math.Pi)
	}
}

func (v *t) Clamp(magnitude float64) T {
	actual := v.Magnitude()
	if actual < magnitude {
		return v
	}
	scale := magnitude / actual
	return v.Scale(scale)
}

func (v *t) Dist(other T) float64 {
	return v.To(other).Magnitude()
}

func (v *t) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func (v *t) Rotate(angle float64) T {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return &t{
		x: v.x*cos - v.y*sin,
		y: v.x*sin + v.y*cos,
	}
}

func (v *t) RotateToward(other T, max float64) T {
	full := v.AngleTo(other)
	if math.Abs(full) <= max {
		return v.Rotate(full)
	} else {
		return v.Rotate(math.Copysign(max, full))
	}
}

func (v *t) Scale(scale float64) T {
	return &t{
		x: v.x * scale,
		y: v.y * scale,
	}
}

func (v *t) To(other T) T {
	return &t{
		x: other.X() - v.x,
		y: other.Y() - v.y,
	}
}

func (v *t) X() float64 {
	return v.x
}

func (v *t) Y() float64 {
	return v.y
}
