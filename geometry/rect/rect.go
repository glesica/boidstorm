package rect

import "github.com/glesica/boidstorm/geometry/vector"

type T interface {
	Center() vector.T
	Height() float64
	LowerLeft() vector.T
	LowerRight() vector.T
	Translate(v vector.T) T
	UpperLeft() vector.T
	UpperRight() vector.T
	Width() float64
}

type t struct {
	lowerLeft, upperRight vector.T
}

func New(xLowerLeft, yLowerLeft, xUpperRight, yUpperRight float64) T {
	return &t{
		vector.New(xLowerLeft, yLowerLeft),
		vector.New(xUpperRight, yUpperRight),
	}
}

func NewCentered(xCenter, yCenter, width, height float64) T {
	xLowerLeft := xCenter - width/2
	yLowerLeft := yCenter - height/2
	xUpperRight := xCenter + width/2
	yUpperRight := yCenter + height/2
	return New(xLowerLeft, yLowerLeft, xUpperRight, yUpperRight)
}

func (r *t) Center() vector.T {
	xDelta := r.UpperRight().X() - r.LowerLeft().X()
	yDelta := r.UpperRight().Y() - r.LowerLeft().Y()
	return vector.New(
		r.LowerLeft().X()+(xDelta/2),
		r.LowerLeft().Y()+(yDelta/2),
	)
}

func (r *t) Height() float64 {
	return r.UpperRight().Y() - r.LowerLeft().Y()
}

func (r *t) LowerLeft() vector.T {
	return r.lowerLeft
}

func (r *t) LowerRight() vector.T {
	return vector.New(
		r.UpperRight().X(),
		r.LowerLeft().Y(),
	)
}

func (r *t) Translate(translation vector.T) T {
	return &t{
		r.lowerLeft.Add(translation),
		r.upperRight.Add(translation),
	}
}

func (r *t) UpperLeft() vector.T {
	return vector.New(
		r.LowerLeft().X(),
		r.UpperRight().Y(),
	)
}

func (r *t) UpperRight() vector.T {
	return r.upperRight
}

func (r *t) Width() float64 {
	return r.UpperRight().X() - r.LowerLeft().X()
}
