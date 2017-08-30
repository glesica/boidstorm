package boid

import "github.com/glesica/boidstorm/geometry/vector"

const DefaultMaxSpeed = 10.0

type Config interface {
	Avoidance() float64
	Conformity() float64
	Exploration() float64
	MaxSpeed() float64
	Wanderlust() (float64, vector.T)

	SetAvoidance(value float64) Config
	SetConformity(value float64) Config
	SetExploration(value float64) Config
	SetMaxSpeed(value float64) Config
	SetWanderlust(value float64, home vector.T) Config
}

type config struct {
	avoidance   float64
	conformity  float64
	exploration float64
	wanderlust  float64
	home        vector.T
	maxSpeed    float64
}

func NewConfig() Config {
	return &config{
		home:     vector.New(0, 0),
		maxSpeed: DefaultMaxSpeed,
	}
}

func (c config) Avoidance() float64 {
	return c.avoidance
}

func (c config) Conformity() float64 {
	return c.conformity
}

func (c config) Exploration() float64 {
	return c.exploration
}

func (c config) Wanderlust() (float64, vector.T) {
	return c.wanderlust, c.home
}

func (c config) MaxSpeed() float64 {
	return c.maxSpeed
}

func (c config) SetAvoidance(value float64) Config {
	c.avoidance = value
	return c
}

func (c config) SetConformity(value float64) Config {
	c.conformity = value
	return c
}

func (c config) SetExploration(value float64) Config {
	c.exploration = value
	return c
}

func (c config) SetWanderlust(value float64, home vector.T) Config {
	c.wanderlust = value
	c.home = home
	return c
}

func (c config) SetMaxSpeed(value float64) Config {
	c.maxSpeed = value
	return c
}
