package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/boidstorm/geometry/circle"
	"github.com/glesica/boidstorm/geometry/line"
	"github.com/glesica/boidstorm/geometry/rect"
	"github.com/glesica/boidstorm/geometry/vector"
	"image/color"
)

type imDrawCanvas struct {
	image  *imdraw.IMDraw
	window *pixelgl.Window
}

func NewIMDrawCanvas(width, height float64) Canvas {
	cfg := pixelgl.WindowConfig{
		Bounds: pixel.Rect{Min: pixel.V(0.0, 0.0), Max: pixel.V(width, height)},
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &imDrawCanvas{
		image:  imdraw.New(nil),
		window: win,
	}
}

func (f *imDrawCanvas) Active() bool {
	return !f.window.Closed()
}

func (f *imDrawCanvas) Circle(shape circle.T, opts DrawOpts) {
	f.image.Color = opts.StrokeColor
	f.image.Push(pixelVec(shape.Center()))
	f.image.Circle(shape.Radius(), opts.StrokeWidth)
	f.image.Reset()
}

func (f *imDrawCanvas) Line(shape line.T, opts DrawOpts) {
	f.image.Color = opts.StrokeColor
	f.image.Push(pixelVec(shape.Start()), pixelVec(shape.End()))
	f.image.Line(opts.StrokeWidth)
	f.image.Reset()
}

func (f *imDrawCanvas) Rect(shape rect.T, opts DrawOpts) {
	f.image.Color = opts.StrokeColor
	f.image.Push(pixelVec(shape.LowerLeft()), pixelVec(shape.UpperLeft()), pixelVec(shape.UpperRight()), pixelVec(shape.LowerRight()))
	f.image.Rectangle(opts.StrokeWidth)
	f.image.Reset()
}

func (f *imDrawCanvas) Size() rect.T {
	bounds := f.window.Bounds()
	return rect.New(
		bounds.Min.X,
		bounds.Min.Y,
		bounds.Max.X,
		bounds.Max.Y,
	)
}

func (f *imDrawCanvas) Update() {
	// TODO: Allow configurable background color
	f.window.Clear(color.Black)
	f.image.Draw(f.window)
	f.image = imdraw.New(nil)
	f.window.Update()
}

func pixelVec(v vector.T) pixel.Vec {
	return pixel.V(v.X(), v.Y())
}
