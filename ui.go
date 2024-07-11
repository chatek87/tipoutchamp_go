package main

import (
	"fmt"

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

	// event loop
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			fmt.Println("frame event received")
			gtx := app.NewContext(&ops, e)
			
			if addBartenderBtn.Clicked(gtx) {
				currentState = BartenderInputView
				fmt.Println(addBartenderBtn.Clicked(gtx)) // Should print true when the button is clicked			
			}

			switch currentState {
			case MainView:
				fmt.Println(currentState)
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
	title := material.H1(th, "TipOut Champ")

	flex := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return title.Layout(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			return layout.Spacer{Height: 20}.Layout(gtx)
		}),
		layout.Rigid(material.Button(th, addBartenderBtn, "Add Bartender").Layout),
	)

	return flex
}

func renderBartenderInputView(gtx layout.Context, th *material.Theme, addBartenderName *widget.Editor, addBartenderHours *widget.Editor, submitButton *widget.Clickable, backButton *widget.Clickable) layout.Dimensions {
	if backButton.Clicked(gtx) {
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

	if submitButton.Clicked(gtx) {
		fmt.Println("Submitting bartender info")
		currentState = MainView // Optionally reset to main view after submission
	} else if backButton.Clicked(gtx) {
		currentState = MainView // Return to main view
	}

	return flex
}
