package gui

import "github.com/glesica/boidstorm/view"

type flowLayout struct {
	widgets []Widget
}

func (f *flowLayout) Count() int {
	panic("implement me")
}

func NewFlowLayout() Layout {
	return &flowLayout{
		widgets: make([]Widget, 0),
	}
}

func (f *flowLayout) Add(widget Widget) {
	panic("implement me")
}

func (f *flowLayout) Draw(canvas view.Canvas) {
	// Maybe look up the algo Java uses?
}

