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
	var staffIcons []widget.Clickable

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
				renderMainView(gtx, th, &addBartenderBtn, staffIcons)
			case BartenderInputView:
				fmt.Println(currentState)
				renderBartenderInputView(gtx, th, &addBartenderName, &addBartenderHours, &submitButton, &backButton)
			}

			e.Frame(gtx.Ops)
		}
	}
}

func renderMainView(gtx C, th *material.Theme, addBartenderBtn *widget.Clickable, staffIcons []widget.Clickable) D {
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
			if len(staffIcons) == 0 {
				for _, b := range calc.BarTeamIn.Bartenders {
					click := widget.Clickable{}
					staffIcons = append(staffIcons, click)
					fmt.Printf("%s\n", b.Name)
				}
			}
			list := layout.List{Axis: layout.Horizontal}
			return list.Layout(gtx, len(staffIcons), func(gtx C, index int) D {
				icon := staffIcons[index]
				return material.Button(th, &icon, "Bartender name").Layout(gtx)
			})
		}),
	)

	return flex
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
