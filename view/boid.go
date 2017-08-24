package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/boidstorm/boid"
	"image/color"
	"github.com/glesica/boidstorm/swarm"
)

// TODO: Draw a boid

type T struct {
	window *pixelgl.Window
	nextFrame *imdraw.IMDraw
}

func New(width, height float64) *T {
	cfg := pixelgl.WindowConfig{
		Bounds: pixel.Rect{pixel.V(0.0, 0.0), pixel.V(width, height)},
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &T{window: win, nextFrame: imdraw.New(nil)}
}

func (v *T) Boid(b *boid.T) {
	v.nextFrame.Push(pixel.V(b.Position().X(), b.Position().Y()))
	v.nextFrame.Circle(10, 1)
}

func (v *T) Swarm(s *swarm.T) {
	s.ForEach(v.Boid)
}

func (v *T) Update() {
	v.window.Clear(color.Black)
	v.nextFrame.Draw(v.window)
	v.window.Update()
	v.nextFrame = imdraw.New(nil)
}

func (v *T) Window() *pixelgl.Window {
	return v.window
}
