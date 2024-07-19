package main

import (
	"image/color"

	"gioui.org/widget"
	"gioui.org/widget/material"
)

type StaffIcon struct {
	Clickable widget.Clickable
	Member    *StaffInput
	Color     color.NRGBA
	Text      string
}

func getStaffIcon(staff StaffInput) StaffIcon {
	v := new(StaffIcon)
	v.Clickable = widget.Clickable{}
	v.Member = &staff
	v.Color = staff.GetPositionColor()
	v.Text = staff.GetFirstInitial()
	return *v
}

func populateStaffIcons(th *material.Theme, calc *Calculator, staffIcons *[]StaffIcon) {
	*staffIcons = []StaffIcon{}
	for _, b := range calc.BarTeamIn.Bartenders {
		s := getStaffIcon(&b)
		*staffIcons = append(*staffIcons, s)
	}
	for _, s := range calc.ServersIn {
		s := getStaffIcon(&s)
		*staffIcons = append(*staffIcons, s)
	}
	for _, e := range calc.EventsIn {
		s := getStaffIcon(&e)
		*staffIcons = append(*staffIcons, s)
	}
	for _, s := range calc.SupportIn {
		s := getStaffIcon(&s)
		*staffIcons = append(*staffIcons, s)
	}
}

func renderStaffIconButtons(th *material.Theme, s *[]StaffIcon) []material.ButtonStyle {
	b := []material.ButtonStyle{}
	for _, icon := range *s {
		btn := material.Button(th, &icon.Clickable, icon.Text)
		btn.Background = icon.Color
		b = append(b, btn)
	}
	return b
}
