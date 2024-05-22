package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// create new window
		w := new(app.Window)
		// app.Config.Title
		w.Option(app.Title("Poop"))
		app.Title("Tipout Chooch")
		app.Size(unit.Dp(400), unit.Dp(600))

		// ops are the operations from the ui
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// th defines the material design style
		th := material.NewTheme()

		// listen for events in the window
		for {
			// grab the event
			evt := w.Event()
			// detect the type
			switch typ := evt.(type) {

			// this is sent when the application should re-render.
			case app.FrameEvent:
				gtx := app.NewContext(&ops, typ)
				title := material.H1(th, "Tipout Chooch")
				title.Color = color.NRGBA{R: 127, G: 0, B: 50, A: 255}
				title.Layout(gtx)
				// Let's try out the flexbox layout:
				layout.Flex{
					// Vertical alignment, from top to bottom
					Axis: layout.Vertical,
					// Empty space is left at the start, i.e. at the top
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					// We insert two rigid elements:
					// First one to hold a button ...
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &startButton, "Start")
							return btn.Layout(gtx)
						},
					),
					// ... then one to hold an empty spacer
					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				typ.Frame(gtx.Ops)
			// and this is sent when the application should exit
			case app.DestroyEvent:
				os.Exit(0)
			}
		}

	}()
	app.Main()
}
