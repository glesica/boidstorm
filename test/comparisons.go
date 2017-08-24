package test

import (
	"github.com/glesica/boidstorm/vector"
	"math"
	"testing"
)

const maxDelta = 0.01

func ExpectEqualFloats(t *testing.T, name string, actual, expected float64) {
	if math.Abs(expected-actual) > maxDelta {
		t.Errorf("%s: %g differs from %g", name, actual, expected)
	}
}

func ExpectEqualVectors(t *testing.T, name string, actual, expected *vector.T) {
	ExpectEqualFloats(t, name+".X()", actual.X(), expected.X())
	ExpectEqualFloats(t, name+".Y()", actual.Y(), expected.Y())
}
