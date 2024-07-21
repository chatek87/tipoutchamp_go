package main

import "image/color"

type Role int

const (
	Bartender Role = iota
	Server
	Event
	Support
)

type StaffInput interface {
	GetFirstInitial() string
	GetPositionColor() color.NRGBA
	GetRole() Role
}

// input models
type BarTeamIn struct {
	Bartenders []BartenderIn
	OwedTo     float64
	Sales      float64
}

// bar
type BartenderIn struct {
	Name  string
	Hours float64
}

func (b *BartenderIn) GetFirstInitial() string {
	fi := string(b.Name[0])
	return fi
}
func (b *BartenderIn) GetPositionColor() color.NRGBA {
	return color.NRGBA{R: 255, G: 0, B: 0, A: 255} // Red for Bartenders
}
func (b *BartenderIn) GetRole() Role {
	return Bartender
}

// server
type ServerIn struct {
	Name   string
	OwedTo float64
	Sales  float64
}

func (s *ServerIn) GetFirstInitial() string {
	fi := string(s.Name[0])
	return fi
}
func (s *ServerIn) GetPositionColor() color.NRGBA {
	return color.NRGBA{G: 255, B: 0, A: 255} // Green for Servers
}
func (s *ServerIn) GetRole() Role {
	return Server
}

// event
type EventIn struct {
	Name    string
	OwedTo  float64
	Sales   float64
	SplitBy int
}

func (e *EventIn) GetFirstInitial() string {
	fi := string(e.Name[0])
	return fi
}
func (e *EventIn) GetPositionColor() color.NRGBA {
	return color.NRGBA{R: 255, G: 165, B: 0, A: 255} // Orange for Events
}
func (e *EventIn) GetRole() Role {
	return Event
}

// support
type SupportIn struct {
	Name  string
	Hours float64
}

func (s *SupportIn) GetFirstInitial() string {
	fi := string(s.Name[0])
	return fi
}
func (s *SupportIn) GetPositionColor() color.NRGBA {
	return color.NRGBA{B: 255, A: 255} // Blue for Support
}
func (s *SupportIn) GetRole() Role {
	return Support
}

// output models ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
type BartenderOut struct {
	Name  string
	Hours float64
	// calculated...
	PercentageOfBarTipPool float64
	OwedToPreTipout        float64
	TipoutToSupport        float64
	TotalAmountTippedOut   float64
	TipoutFromServers      float64
	TipoutFromEvents       float64
	TotalTipoutReceived    float64
	FinalPayout            float64
}

type BarTeamOut struct {
	Bartenders      []BartenderOut
	OwedToPreTipout float64
	Sales           float64
	// calculated...
	TipoutToSupport      float64
	TotalAmountTippedOut float64
	TipoutFromServers    float64
	TipoutFromEvents     float64
	TotalTipoutReceived  float64
	FinalPayout          float64
}

type ServerOut struct {
	Name            string
	OwedToPreTipout float64
	Sales           float64
	// calculated...
	TipoutToBar          float64
	TipoutToSupport      float64
	TotalAmountTippedOut float64
	FinalPayout          float64
}

type EventOut struct {
	Name            string
	OwedToPreTipout float64
	Sales           float64
	SplitBy         int // the number of workers splitting the final payout
	// calculated...
	TipoutToBar          float64
	TipoutToSupport      float64
	TotalAmountTippedOut float64
	FinalPayout          float64
	FinalPayoutPerWorker float64
}

type SupportOut struct {
	Name  string
	Hours float64
	// calculated...
	PercentageOfSupportTipPool float64
	TipoutFromBar              float64
	TipoutFromServers          float64
	TipoutFromEvents           float64
	FinalPayout                float64
}
