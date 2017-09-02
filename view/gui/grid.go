package gui

import (
	"github.com/glesica/boidstorm/view"
)

// We may need to export GridLayout so the consumer can specify
// the number of rows and columns. Or we could export a config
// type that gets passed to the New* method.

type gridLayout struct {
	widgets []Widget
}

// func NewGridLayout(config GridLayoutConfig) Layout {
func NewGridLayout() Layout {
	return &gridLayout{
		widgets: make([]Widget, 0),
	}
}

func (g *gridLayout) Draw(canvas view.Canvas) {
	count := len(g.widgets)
	// compute number of rows / columns (sqrt)
	// draw each widget
}

func (g *gridLayout) Add(widget Widget) {
	panic("implement me")
}
