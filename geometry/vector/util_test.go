package vector_test

import (
	"testing"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/test"
)

func TestCentroid(t *testing.T) {
	v0 := vector.New(1, 0)
	v1 := vector.New(-1, 0)
	v2 := vector.New(0,1)
	v3 := vector.New(0,-1)
	c := vector.Centroid(v0, v1, v2, v3)
	test.ExpectEqualVectors(t, "c=(0,0)", c, vector.Zero())
}
