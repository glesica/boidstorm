package vector_test

import (
	"github.com/glesica/boidstorm/test"
	"github.com/glesica/boidstorm/vector"
	"math"
	"testing"
)

func TestNew(t *testing.T) {
	v0 := vector.New(0.0, 0.0)
	test.ExpectEqualFloats(t, "v0.X()", v0.X(), 0.0)
	test.ExpectEqualFloats(t, "v0.Y()", v0.Y(), 0.0)

	v1 := vector.New(1.0, -1.0)
	test.ExpectEqualFloats(t, "v1.X()", v1.X(), 1.0)
	test.ExpectEqualFloats(t, "v1.Y()", v1.Y(), -1.0)
}

func TestZero(t *testing.T) {
	v0 := vector.Zero()
	test.ExpectEqualFloats(t, "v0.X()", v0.X(), 0.0)
	test.ExpectEqualFloats(t, "v0.Y()", v0.Y(), 0.0)
}

func TestT_Add(t *testing.T) {
	v0 := vector.New(0.0, 0.0)
	v1 := vector.New(1.0, 2.0)
	test.ExpectEqualVectors(t, "v0.Add(v1)", v0.Add(v1), v1)
}

func TestT_Dist(t *testing.T) {
	v0 := vector.Zero()
	v1 := vector.New(1.0, 0.0)
	test.ExpectEqualFloats(t, "v0.Dist(v1)", v0.Dist(v1), 1.0)
}

func TestT_Magnitude(t *testing.T) {
	v0 := vector.New(3.0, 4.0)
	test.ExpectEqualFloats(t, "v0.Magnitude()", v0.Magnitude(), 5.0)

	v1 := vector.New(-3.0, -4.0)
	test.ExpectEqualFloats(t, "v1.Magnitude()", v1.Magnitude(), 5.0)
}

func TestT_Rotate(t *testing.T) {
	v0 := vector.New(1.0, 0.0)
	v1 := vector.New(-1.0, 0.0)
	test.ExpectEqualVectors(t, "v0.Rotate(Pi)", v0.Rotate(math.Pi), v1)
	test.ExpectEqualVectors(t, "v1.Rotate(-Pi)", v1.Rotate(-math.Pi), v0)
}

func TestT_Scale(t *testing.T) {
	v0 := vector.New(1.0, 1.0)
	v1 := vector.New(0.5, 0.5)
	test.ExpectEqualVectors(t, "v0.Scale(0.5)", v0.Scale(0.5), v1)
	test.ExpectEqualVectors(t, "v1.Scale(2.0)", v1.Scale(2.0), v0)
}
