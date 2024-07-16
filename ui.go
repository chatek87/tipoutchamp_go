package main

import (
	"fmt"
	"strconv"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func loop(w *app.Window) error {
	var ops op.Ops
	th := material.NewTheme()

	// UI elements
	var addBartenderBtn widget.Clickable
	var addBartenderName widget.Editor
	var addBartenderHours widget.Editor
	var submitButton widget.Clickable
	var backButton widget.Clickable

	// event loop
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			fmt.Printf("frame event received. viewstate: %d\n", currentState)
			gtx := app.NewContext(&ops, e)

			if addBartenderBtn.Clicked(gtx) {
				currentState = BartenderInputView
				fmt.Println(addBartenderBtn.Clicked(gtx)) // Should print true when the button is clicked
			}

			switch currentState {
			case MainView:
				renderMainView(gtx, th, &addBartenderBtn)
			case BartenderInputView:
				fmt.Println(currentState)
				renderBartenderInputView(gtx, th, &addBartenderName, &addBartenderHours, &submitButton, &backButton)
			}

			e.Frame(gtx.Ops)
		}
	}
}

func renderMainView(gtx layout.Context, th *material.Theme, addBartenderBtn *widget.Clickable) layout.Dimensions {	
	flex := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			title := material.H1(th, "TipOut Champ")
			return title.Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return layout.Spacer{Height: 20}.Layout(gtx)
		}),
		layout.Rigid(material.Button(th, addBartenderBtn, "Add Bartender").Layout),
		layout.Rigid(func(gtx C) D {
			// TODO: modify this func to print all the details of the staff input models w/ a clickable icon for editing
			// for _, bartender := range calc.BarTeamIn.Bartenders {
			// 	drawStaffIcon(bartender.Name, Bartender)
			// }
			// for _, server := range calc.ServersIn {
			// 	drawStaffIcon(server.Name, Server)

			// }
			// for _, event := range calc.EventsIn {
			// 	drawStaffIcon(event.Name, Event)

			// }
			// for _, support := range calc.SupportIn {
			// 	drawStaffIcon(support.Name, Support)

			// }
			return material.Label(th, unit.Sp(12), "poop").Layout(gtx)
		}),
	)

	return flex
}

func renderBartenderInputView(gtx layout.Context, th *material.Theme, addBartenderName *widget.Editor, addBartenderHours *widget.Editor, submitButton *widget.Clickable, backButton *widget.Clickable) layout.Dimensions {
	if backButton.Clicked(gtx) {
		currentState = MainView
	}

	if submitButton.Clicked(gtx) {
		fmt.Println("Submitting bartender info")
		b := new(BartenderIn)
		b.Name = addBartenderName.Text()
		hoursText := addBartenderHours.Text()
		
		hours, err := strconv.ParseFloat(hoursText, 64)
		if err != nil {
			fmt.Printf("Error converting hours to float: %v\n", err)
			// return err
		}
		b.Hours = hours
		// add b to bartenders
		calc.BarTeamIn.Bartenders = append(calc.BarTeamIn.Bartenders, *b)
		for _, b := range calc.BarTeamIn.Bartenders {
			fmt.Printf(b.Name)
		}
		currentState = MainView 
		} else if backButton.Clicked(gtx) {
			currentState = MainView 
		}
	flex := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return material.Editor(th, addBartenderName, "Bartender Name").Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return material.Editor(th, addBartenderHours, "Bartender Hours").Layout(gtx)
		}),
		layout.Rigid(material.Button(th, submitButton, "Submit").Layout),
		layout.Rigid(material.Button(th, backButton, "Go Back").Layout),
	)		
		return flex
}
