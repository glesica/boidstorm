package boid

import "github.com/glesica/boidstorm/geometry/vector"

const DefaultMaxSpeed = 1.0

type Config interface {
	Avoidance() float64
	Conformity() float64
	Exploration() float64

	// TODO: Need to include a modifier as well as the position
	Home() vector.T
	MaxSpeed() float64

	SetAvoidance(value float64) Config
	SetConformity(value float64) Config
	SetExploration(value float64) Config
	SetHome(value vector.T) Config
	SetMaxSpeed(value float64) Config
}

type config struct {
	avoidance  float64
	conformity float64
	exploration float64
	home vector.T
	maxSpeed float64
}

func NewConfig() Config {
	return &config{
		home: vector.New(0, 0),
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

func (c config) Home() vector.T {
	return c.home
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

func (c config) SetHome(value vector.T) Config {
	c.home = value
	return c
}

func (c config) SetMaxSpeed(value float64) Config {
	c.maxSpeed = value
	return c
}
