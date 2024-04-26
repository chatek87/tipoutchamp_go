package main

// input models
type BartenderIn struct {
	Name  string
	Hours float64
}

type BarTeamIn struct {
	Bartenders []BartenderIn
	OwedTo     float64
	Sales      float64
}

type ServerIn struct {
	Name   string
	OwedTo float64
	Sales  float64
}

type EventIn struct {
	Name    string
	OwedTo  float64
	Sales   float64
	SplitBy int
}

type SupportIn struct {
	Name  string
	Hours float64
}

// output models ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
type BartenderOut struct {
	Name                 string
	Hours                float64
	OwedToPreTipout      float64
	TipoutToSupport      float64
	TotalAmountTippedOut float64
	FinalPayout          float64
}

type BarTeamOut struct {
	Bartenders           []BartenderOut
	OwedToPreTipout      float64
	Sales                float64
	TipoutToSupport      float64
	TotalAmountTippedOut float64
	FinalPayout          float64
}

type ServerOut struct {
	Name                 string
	OwedToPreTipout      float64
	Sales                float64
	TipoutToBar          float64
	TipoutToSupport      float64
	TotalAmountTippedOut float64
	FinalPayout          float64
}

type EventOut struct {
	Name                 string
	OwedToPreTipout      float64
	Sales                float64
	TipoutToBar          float64
	TipoutToSupport      float64
	TotalAmountTippedOut float64
	FinalPayout          float64
	SplitBy              int
}

type SupportOut struct {
	Name              string
	Hours             float64
	TipoutFromBar     float64
	TipoutFromServers float64
	TipoutFromEvents  float64
	FinalPayout       float64
}
