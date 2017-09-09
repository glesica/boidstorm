package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/boidstorm/view/imdraw"
	"github.com/glesica/boidstorm/view/gui"
)

func run() {
	f := imdraw.NewIMDrawCanvas(1024, 768)
	l := gui.NewLinearLayout(false)
	b0 := gui.NewButton()
	l.Add(b0)
	b1 := gui.NewButton()
	l.Add(b1)

	for f.Active() {
		l.Draw(f)
		f.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
