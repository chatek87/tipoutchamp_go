package main

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type StaffIcon struct {
	Clickable widget.Clickable
	Member    *StaffInput
}

func GetStaffIcon(staff StaffInput) StaffIcon {
	v := new(StaffIcon)
	v.Clickable = widget.Clickable{}
	v.Member = &staff
	return *v
}

func populateStaffIcons(th *material.Theme, calc *Calculator, staffIcons *[]StaffIcon) {
	for _, b := range calc.BarTeamIn.Bartenders {
		click := GetStaffIcon(&b) // Directly passing the address of b, assuming BartenderIn implements StaffInput
		*staffIcons = append(*staffIcons, click)
	}
	for _, s := range calc.ServersIn {
		click := GetStaffIcon(&s) // Directly passing the address of s, assuming ServerIn implements StaffInput
		*staffIcons = append(*staffIcons, click)
	}
	for _, e := range calc.EventsIn {
		click := GetStaffIcon(&e) // Directly passing the address of e, assuming EventIn implements StaffInput
		*staffIcons = append(*staffIcons, click)
	}
	for _, s := range calc.SupportIn {
		click := GetStaffIcon(&s) // Directly passing the address of s, assuming SupportIn implements StaffInput
		*staffIcons = append(*staffIcons, click)
	}
}
