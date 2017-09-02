package gui

type Button interface {
	Widget

	// OnClick sets the click handler for this button.
	OnClick(func())
}
