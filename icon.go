package main

import (
	"image/color"
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

func getPositionColor(p Position) color.RGBA {
	switch p {
	case Bartender:
		return color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red for Bartenders
	case Server:
		return color.RGBA{G: 255, B: 0, A: 255} // Green for Servers
	case Event:
		return color.RGBA{R: 255, G: 165, B: 0, A: 255} // Orange for Events
	case Support:
		return color.RGBA{B: 255, A: 255} // Blue for Support
	default:
		return color.RGBA{A: 255} // Transparent for unknown positions
	}
}
