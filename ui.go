package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions


func loop(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	// UI elements
	var addBartenderBtn widget.Clickable
	var addBartenderName widget.Editor
	var addBartenderHours widget.Editor
	var addBartenderBtnSelected bool
	// event loop
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			title := material.H1(th, "hey bro")

			if addBartenderBtn.Clicked(gtx) {
				addBartenderBtnSelected = !addBartenderBtnSelected
				// do something, like have a few editors appear along with a 'add' button and a 'go back' or 'clear' or something
			}

			layout.W.Layout(gtx, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						return title.Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						return layout.Spacer{Height: 20}.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: 200}.Layout),
					layout.Rigid(func(gtx C) D {
						return material.Button(th, &addBartenderBtn, "add bartender").Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						if addBartenderBtnSelected {
							flex := layout.Flex{Axis: layout.Vertical}
							return flex.Layout(gtx,
								layout.Rigid(func(gtx C) D {
									return material.Editor(th, &addBartenderName, "bartender name").Layout(gtx)
								}),
								layout.Rigid(func(gtx C) D {
									return material.Editor(th, &addBartenderHours, "bartender hours").Layout(gtx)
								}))
						} else {
							return material.H4(th, fmt.Sprintf(": /")).Layout(gtx)
						}
					}),
					layout.Rigid(func(gtx C) D {
						return material.H6(th, fmt.Sprintf("here's some filler text")).Layout(gtx)
					}),
				)
			})
			// Pass the drawing operations to the GPU. (send the operations from the context gtx to the FrameEvent)
			e.Frame(gtx.Ops)
		}
	}
}

// func bartenderInputScreen(w *app.Window) error {

// }
