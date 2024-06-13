package main

import (
	"errors"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

var addBartenderInput widget.Editor
var showBartenderForm bool

func main() {
	go func() {
		// create new window
		w := new(app.Window)
		w.Option(app.Title("tipout champ"), app.Size(unit.Dp(400), unit.Dp(600)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	// ops are the OPERATIONS from the UI
	var ops op.Ops

	// var addBartenderInput widget.Editor

	th := material.NewTheme()

	// listen for events in the window 	(this is the EVENT LOOP)
	for {
		// first grab the event
		evt := w.Event()

		// then detect the type  (this is a TYPE SWITCH)
		switch typ := evt.(type) {

		// this is sent when the app should re-render.
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ) // define a new GRAPHICAL CONTEXT (gtx)

			// if addBartenderBtn.Clicked(gtx) {
			// 	showBartenderForm = !showBartenderForm
			// 	// render editors
			// }

			layout.Flex{
				// vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// the inputbox
				layout.Rigid(
					func(gtx C) D {
						// Wrap the editor in material design
						ed := material.Editor(th, &addBartenderInput, "bartender name")

						// Define characteristics of the input box
						// addBartenderInput.SingleLine = true
						addBartenderInput.Alignment = text.Middle

						// Define insets ...
						margins := layout.UniformInset(10)
						// margins := layout.Inset{
						// 	Top:    unit.Dp(0),
						// 	Right:  unit.Dp(170),
						// 	Bottom: unit.Dp(40),
						// 	Left:   unit.Dp(170),
						// }
						// ... and borders ...
						border := widget.Border{
							Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
							CornerRadius: unit.Dp(3),
							Width:        unit.Dp(3),
						}
						// ... before laying it out, one inside the other
						return margins.Layout(gtx,
							func(gtx C) D {
								return border.Layout(gtx, ed.Layout)
							},
						)
					},
				),
				// the button
				// layout.Rigid( // Rigid() accepts a WIDGET. A widget is simply something that returns its own DIMENSIONS.
				// 	func(gtx C) D {
				// 		// ONE: First define margins around the button using layout.Inset ...
				// 		margins := layout.Inset{
				// 			Top:    unit.Dp(25),
				// 			Bottom: unit.Dp(25),
				// 			Left:   unit.Dp(35),
				// 			Right:  unit.Dp(35),
				// 		}
				// 		// marginsAutoSpacedEvenly := layout.UniformInset(25)	// we can also do this!
				// 		// TWO: ... then we lay out those margins ...
				// 		return margins.Layout(gtx,
				// 			// THREE: ... and finally within the margins, we define and lay out the button
				// 			func(gtx C) D {
				// 				var text string
				// 				text = "Start"
				// 				if boiling && progress < 1 {
				// 					text = "Stop"
				// 				}
				// 				if boiling && progress >= 1 {
				// 					text = "Finished"
				// 				}
				// 				btn := material.Button(th, &startButton, text) // define button
				// 				return btn.Layout(gtx)
				// 			},
				// 		)
				// 	},
				// ),
			)
			typ.Frame(gtx.Ops) // send the operations from the context gtx to the FrameEvent

		// and this is sent when the app should exit
		case app.DestroyEvent:
			return errors.New("user exited the application")
		}
	}
}
