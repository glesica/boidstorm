package gui

import "github.com/glesica/boidstorm/view"

type Layout interface {
	view.Drawer

	Add(widget Widget)
}
