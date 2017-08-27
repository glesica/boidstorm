package boid_test

import (
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/test"
	"testing"
)

func TestConfig_Avoidance(t *testing.T) {
	c0 := boid.NewConfig()
	c1 := c0.SetAvoidance(0.5)

	test.ExpectEqualFloats(t, "c0.Avoidance()", c0.Avoidance(), 0.0)
	test.ExpectEqualFloats(t, "c1.Avoidance()", c1.Avoidance(), 0.5)
}

func TestConfig_Conformity(t *testing.T) {
	c0 := boid.NewConfig()
	c1 := c0.SetConformity(0.5)

	test.ExpectEqualFloats(t, "c0.Conformity()", c0.Conformity(), 0.0)
	test.ExpectEqualFloats(t, "c1.Conformity()", c1.Conformity(), 0.5)
}

func TestConfig_Exploration(t *testing.T) {
	c0 := boid.NewConfig()
	c1 := c0.SetExploration(0.5)

	test.ExpectEqualFloats(t, "c0.Exploration()", c0.Exploration(), 0.0)
	test.ExpectEqualFloats(t, "c1.Exploration()", c1.Exploration(), 0.5)
}

func TestConfig_MaxSpeed(t *testing.T) {
	c0 := boid.NewConfig()
	test.ExpectEqualFloats(t, "default MaxSpeed", c0.MaxSpeed(), boid.DefaultMaxSpeed)
	c1 := c0.SetMaxSpeed(0.5)

	test.ExpectEqualFloats(t, "c0.MaxSpeed()", c0.MaxSpeed(), 0.0)
	test.ExpectEqualFloats(t, "c1.MaxSpeed()", c1.MaxSpeed(), 0.5)
}
