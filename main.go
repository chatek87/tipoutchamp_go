package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("tipout champ"), app.Size(unit.Dp(400), unit.Dp(600)))
		// alternate syntax:
		// err := draw(w)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	// UI elements
	var addBartenderBtn widget.Clickable
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
						return material.Button(th, &addBartenderBtn, "add bartender").Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						if addBartenderBtnSelected {
							return material.H4(th, fmt.Sprintf("SELECTED")).Layout(gtx)
						} else {
							return material.H4(th, fmt.Sprintf(": /")).Layout(gtx)
						}
					}),
				)
			})
			// Pass the drawing operations to the GPU. (send the operations from the context gtx to the FrameEvent)
			e.Frame(gtx.Ops)
		}
	}
}
