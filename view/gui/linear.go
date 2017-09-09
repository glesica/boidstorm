package gui

import (
	"github.com/glesica/boidstorm/view"
	"github.com/glesica/boidstorm/geometry/rect"
)

type linearLayout struct {
	isColumn bool
	widgets []Widget
}

func NewLinearLayout(isColumn bool) Layout {
	return &linearLayout{
		isColumn: isColumn,
		widgets: make([]Widget, 0),
	}
}

func (l *linearLayout) Add(widget Widget) {
	l.widgets = append(l.widgets, widget)
}

func (l *linearLayout) Count() int {
	return len(l.widgets)
}

func (l *linearLayout) Draw(canvas view.Canvas) {
	count := float64(l.Count())

	width := canvas.Size().Width()
	height := canvas.Size().Height()

	unitWidth := width / count
	unitHeight := height / count

	for i, widget := range l.widgets {
		offset := float64(i)

		var startX float64
		var startY float64
		var endX float64
		var endY float64

		if l.isColumn {
			startX = 0.0
			startY = offset * unitHeight
			endX = width
			endY = unitHeight + offset * unitHeight
		} else {
			startX = offset * unitWidth
			startY = 0.0
			endX = unitWidth + offset * unitWidth
			endY = height
		}

		widget.Draw(canvas.Region(rect.New(
			startX,
				startY,
					endX,
						endY,
		)))
	}
}
