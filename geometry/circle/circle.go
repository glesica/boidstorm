package circle

import "github.com/glesica/boidstorm/geometry/vector"

type T interface {
	Center() vector.T
	Height() float64
	Radius() float64
	Translate(v vector.T) T
	Width() float64
}

type t struct {
	center vector.T
	radius float64
}

func New(xCenter, yCenter, radius float64) T {
	return &t{
		center: vector.New(xCenter, yCenter),
		radius: radius,
	}
}

func (c *t) Center() vector.T {
	return c.center
}

func (c *t) Height() float64 {
	return 2 * c.radius
}

func (c *t) Radius() float64 {
	return c.radius
}

func (c *t) Translate(v vector.T) T {
	panic("implement me")
}

func (c *t) Width() float64 {
	return 2 * c.radius
}
