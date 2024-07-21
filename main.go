package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
)

type C = layout.Context
type D = layout.Dimensions

type ViewState int

const (
	MainView ViewState = iota
	BartenderInputView
	StaffMemberDetailView
)

var currentState ViewState = MainView
var calc *Calculator

func main() {
	calc = new(Calculator)
	calc.SeedSampleData()
	calc.RunCalculationsPopulateOutputFields()

	go func() {
		w := new(app.Window)
		w.Option(app.Title("TipOut Champ"))

		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
