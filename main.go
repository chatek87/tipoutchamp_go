package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
)

var calc *Calculator

func init() {
	calc = new(Calculator)
}

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("tipout champ"), app.Size(unit.Dp(400), unit.Dp(600)))

		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
