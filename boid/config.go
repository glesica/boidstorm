package boid

type Config interface {
	Avoidance() float64
	Conformity() float64
	MaxSpeed() float64
	SetAvoidance(value float64) Config
	SetConformity(value float64) Config
	SetMaxSpeed(value float64) Config
}

type config struct {
	avoidance  float64
	conformity float64
	maxSpeed float64
}

func NewConfig() Config {
	return &config{
		maxSpeed: 1.0,
	}
}

func (c config) Avoidance() float64 {
	return c.avoidance
}

func (c config) Conformity() float64 {
	return c.conformity
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

func (c config) SetMaxSpeed(value float64) Config {
	c.maxSpeed = value
	return c
}
