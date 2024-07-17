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

func GenerateClickableForStaffMember(th *material.Theme, staff StaffInput) widget.Clickable {
	click := widget.Clickable{}
	button := material.Button(th, &click, staff.GetFirstInitial())
	button.Background = staff.GetPositionColor()
	return click
}

func renderStaffIcons(gtx C, th *material.Theme, calc *Calculator, staffIcons []widget.Clickable) []widget.Clickable {
	// var staffIcons []widget.Clickable

	for _, b := range calc.BarTeamIn.Bartenders {
		click := GenerateClickableForStaffMember(th, &b)
		staffIcons = append(staffIcons, click)
	}
	for _, s := range calc.ServersIn {
		click := GenerateClickableForStaffMember(th, &s)
		staffIcons = append(staffIcons, click)
	}
	for _, e := range calc.EventsIn {
		click := GenerateClickableForStaffMember(th, &e)
		staffIcons = append(staffIcons, click)
	}
	for _, s := range calc.SupportIn {
		click := GenerateClickableForStaffMember(th, &s)
		staffIcons = append(staffIcons, click)
	}
	return staffIcons
}
