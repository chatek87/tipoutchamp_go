package main

import (
	"fmt"
	"strconv"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
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
	var staffIcons []StaffIcon

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

			for _, member := range staffIcons {
				if member.Clickable.Clicked(gtx) {
					currentState = StaffMemberDetailView
					renderStaffDetailView(member.Role)
				}
			}

			switch currentState {
			case MainView:
				renderMainView(gtx, th, &addBartenderBtn, &staffIcons)
			case BartenderInputView:
				fmt.Println(currentState)
				renderBartenderInputView(gtx, th, &addBartenderName, &addBartenderHours, &submitButton, &backButton)
			}

			e.Frame(gtx.Ops)
		}
	}
}

func renderMainView(gtx C, th *material.Theme, addBartenderBtn *widget.Clickable, staffIcons *[]StaffIcon) D {
	flex := layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceEvenly,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx C) D {
				title := material.H1(th, "TipOut Champ")
				return title.Layout(gtx)
			}),
		layout.Rigid(func(gtx C) D {
			return layout.Spacer{Height: 20}.Layout(gtx)
		}),
		layout.Rigid(material.Button(th, addBartenderBtn, "Add Bartender").Layout),
		layout.Rigid(func(gtx C) D {
			// here we need to 1) clear staffIcons slice --> handled in populateStaffIcons()
			// 2) populate staffIcons based on calc contents --> handled in populateStaffIcons()
			populateStaffIcons(calc, staffIcons)
			// 3) use the clickables we have generated to draw the buttons (use material.Button() to return a []material.ButtonStyle)
			iconButtons := renderStaffIconButtons(th, staffIcons)
			// 4) somehow pass an identifier (either pointer or index) of each element so that when we click a button, it knows which element of the slice is associated w/ that button

			// 5) wrap all of these buttons in a layout.List and lay them out
			list := layout.List{Axis: layout.Horizontal}
			return list.Layout(gtx, len(iconButtons)*2-1, func(gtx C, index int) D {
				if index%2 == 1 {
					return layout.Spacer{Width: 10}.Layout(gtx) // Adjust the width as needed
				}
				return iconButtons[index/2].Layout(gtx)
			})
		}),
	)

	return flex
}

func renderStaffDetailView(role Role) {
	//TODO: render each w/ editors for the appropriate fields
	switch role {
	case Bartender:
	case Server:
	case Event:
	case Support:
	}
}

func renderBartenderInputView(gtx C, th *material.Theme, addBartenderName *widget.Editor, addBartenderHours *widget.Editor, submitButton *widget.Clickable, backButton *widget.Clickable) D {
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
