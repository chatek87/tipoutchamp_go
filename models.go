package main

type BarTeam struct {
	Bartenders   []Bartender
	BarTeamSales float64
}

type Bartender struct {
	Name  string
	Hours float64
}

type Server struct {
	Name   string
	OwedTo float64
	Sales  float64
}

type Event struct {
	Name    string
	OwedTo  float64
	Sales   float64
	SplitBy int
}

type Support struct {
	Name  string
	Hours float64
}
