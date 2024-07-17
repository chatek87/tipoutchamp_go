package main

import (
	"image/color"

	"gioui.org/widget"
	"gioui.org/widget/material"
)

// i feel like encapsulating this stuff is unnecessary since we are
// passing this info from the parent widget...(?)
type StaffIcon struct {
	NameInitial string
	Position    string
	Color       color.RGBA
}

type Position int

const (
	Bartender Position = iota
	Server
	Event
	Support
)

func drawStaffIcon(th *material.Theme, name string, p Position) widget.Clickable {
	click := widget.Clickable{}
	button := material.Button(th, &click, name)
	button.Background = getPositionColor(p)
	return click
}

func getPositionColor(p Position) color.NRGBA {
	switch p {
	case Bartender:
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255} // Red for Bartenders
	case Server:
		return color.NRGBA{G: 255, B: 0, A: 255} // Green for Servers
	case Event:
		return color.NRGBA{R: 255, G: 165, B: 0, A: 255} // Orange for Events
	case Support:
		return color.NRGBA{B: 255, A: 255} // Blue for Support
	default:
		return color.NRGBA{A: 255} // Transparent for unknown positions
	}
}
