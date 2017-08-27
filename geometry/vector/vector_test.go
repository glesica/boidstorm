package vector_test

import (
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/test"
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

func TestT_Angle(t *testing.T) {
	v0 := vector.New(1, 1)
	test.ExpectEqualFloats(t, "v0.Angle()", v0.Angle(), math.Pi/4)

	v1 := vector.New(-1, 1)
	test.ExpectEqualFloats(t, "v1.Angle()", v1.Angle(), math.Pi*3/4)

	v2 := vector.New(-1, -1)
	test.ExpectEqualFloats(t, "v2.Angle()", v2.Angle(), -math.Pi*3/4)

	v3 := vector.New(1, -1)
	test.ExpectEqualFloats(t, "v3.Angle()", v3.Angle(), -math.Pi/4)
}

func TestT_AngleTo(t *testing.T) {
	v0 := vector.New(1, 1)
	v1 := vector.New(-1,1)
	v2 := vector.New(-1,-1)
	v3 := vector.New(1,-1)

	test.ExpectEqualFloats(t, "v0.AngleTo(v1)", v0.AngleTo(v1), math.Pi / 2)
	test.ExpectEqualFloats(t, "v0.AngleTo(v3)", v0.AngleTo(v3), -math.Pi / 2)

	test.ExpectEqualFloats(t, "v1.AngleTo(v2)", v1.AngleTo(v2), math.Pi / 2)
	test.ExpectEqualFloats(t, "v1.AngleTo(v0)", v1.AngleTo(v0), -math.Pi / 2)

	test.ExpectEqualFloats(t, "v2.AngleTo(v3)", v2.AngleTo(v3), math.Pi / 2)
	test.ExpectEqualFloats(t, "v2.AngleTo(v1)", v2.AngleTo(v1), -math.Pi / 2)

	test.ExpectEqualFloats(t, "v3.AngleTo(v0)", v3.AngleTo(v0), math.Pi / 2)
	test.ExpectEqualFloats(t, "v3.AngleTo(v2)", v3.AngleTo(v2), -math.Pi / 2)
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
	v0 := vector.New(1, 0)
	v1 := vector.New(-1, 0)

	test.ExpectEqualVectors(t, "v0.Rotate(Pi)", v0.Rotate(math.Pi), v1)
	test.ExpectEqualVectors(t, "v1.Rotate(-Pi)", v1.Rotate(-math.Pi), v0)

	v2 := vector.New(1, 1)
	v3 := vector.New(-1, 1)
	v4 := vector.New(-1, -1)
	v5 := vector.New(1, -1)

	test.ExpectEqualVectors(t, "v2.Rotate(Pi/2)", v2.Rotate(math.Pi/2), v3)
	test.ExpectEqualVectors(t, "v2.Rotate(-Pi/2)", v2.Rotate(-math.Pi/2), v5)

	test.ExpectEqualVectors(t, "v3.Rotate(Pi/2)", v3.Rotate(math.Pi/2), v4)
	test.ExpectEqualVectors(t, "v3.Rotate(-Pi/2)", v3.Rotate(-math.Pi/2), v2)

	test.ExpectEqualVectors(t, "v4.Rotate(Pi/2)", v4.Rotate(math.Pi/2), v5)
	test.ExpectEqualVectors(t, "v4.Rotate(-Pi/2)", v4.Rotate(-math.Pi/2), v3)

	test.ExpectEqualVectors(t, "v5.Rotate(Pi/2)", v5.Rotate(math.Pi/2), v2)
	test.ExpectEqualVectors(t, "v5.Rotate(-Pi/2)", v5.Rotate(-math.Pi/2), v4)
}

func TestT_Scale(t *testing.T) {
	v0 := vector.New(1.0, 1.0)
	v1 := vector.New(0.5, 0.5)
	test.ExpectEqualVectors(t, "v0.Scale(0.5)", v0.Scale(0.5), v1)
	test.ExpectEqualVectors(t, "v1.Scale(2.0)", v1.Scale(2.0), v0)
}
