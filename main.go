package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
)

type C = layout.Context
type D = layout.Dimensions

type ViewState int	

const (
	MainView ViewState = iota
	BartenderInputView
)

var currentState ViewState = MainView
var calc *Calculator

func main() {
	calc = new(Calculator)
	calc.SeedSampleData()

	go func() {
		w := new(app.Window)
		w.Option(app.Title("TipOut Champ"), app.Size(unit.Dp(400), unit.Dp(600)))

		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
