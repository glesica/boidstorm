package gui

import (
	"github.com/glesica/boidstorm/view"
	"github.com/glesica/boidstorm/geometry/rect"
	"golang.org/x/image/colornames"
	"github.com/glesica/boidstorm/geometry/line"
	"github.com/glesica/boidstorm/geometry/vector"
)

type Button interface {
	Widget

	// OnClick sets the click handler for this button.
	OnClick(func())
}

type button struct {}

func NewButton() Button {
	return &button{}
}

func (b *button) Draw(canvas view.Canvas) {
	stroke := 4.0

	lightOpts := view.DrawOpts{
		StrokeWidth: stroke,
		StrokeColor: colornames.White,
	}
	darkOpts := view.DrawOpts{
		StrokeWidth: stroke,
		StrokeColor: colornames.Darkgray,
	}

	size := canvas.Size()
	center := size.Center()
	margin := 30.0
	width := size.Width() - 2.0 * margin
	height := 90.0
	gap := 14.0
	innerWidth := width - 2 * gap
	innerHeight := height - 2 * gap

	lowerLeft := center.Add(vector.New(-width / 2.0, -height / 2.0))
	lowerRight := center.Add(vector.New(width / 2.0, -height / 2.0))
	upperRight := center.Add(vector.New(width / 2.0, height / 2.0))
	upperLeft := center.Add(vector.New(-width/2.0, height/2.0))

	innerLowerLeft := center.Add(vector.New(-innerWidth / 2.0, -innerHeight / 2.0))
	innerLowerRight := center.Add(vector.New(innerWidth / 2.0, -innerHeight / 2.0))
	innerUpperRight := center.Add(vector.New(innerWidth / 2.0, innerHeight / 2.0))
	innerUpperLeft := center.Add(vector.New(-innerWidth / 2.0, innerHeight / 2.0))

	canvas.Line(line.New(lowerLeft, lowerRight), lightOpts)
	canvas.Line(line.New(lowerRight, upperRight), lightOpts)
	canvas.Line(line.New(upperRight, upperLeft), darkOpts)
	canvas.Line(line.New(upperLeft, lowerLeft), darkOpts)

	canvas.Line(line.New(innerLowerLeft, innerLowerRight), darkOpts)
	canvas.Line(line.New(innerLowerRight, innerUpperRight), darkOpts)
	canvas.Line(line.New(innerUpperRight, innerUpperLeft), lightOpts)
	canvas.Line(line.New(innerUpperLeft, innerLowerLeft), lightOpts)
}

func (b *button) MaxSize() rect.T {
	panic("implement me")
}

func (b *button) MinSize() rect.T {
	panic("implement me")
}

func (b *button) OnClick(func()) {
	panic("implement me")
}



