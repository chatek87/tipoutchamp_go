package main

type Calculator struct {
	// input fields
	BarTeamIn BarTeamIn
	ServersIn []ServerIn
	EventsIn  []EventIn
	SupportIn []SupportIn

	// output fields
	BarTeamOut BarTeamOut
	ServersOut []ServerOut
	EventsOut  []EventOut
	SupportOut []SupportOut

	// tip pool info fields needed for calculations
	BarPool             float64
	TotalBarHours       float64
	BarCount            int
	BarTipoutPercentage float64

	SupportPool             float64
	TotalSupportHours       float64
	SupportCount            int
	SupportTipoutPercentage float64

	ToSupportFromBarAmount     float64
	ToSupportFromServersAmount float64
	ToSupportFromEventsAmount  float64
}

func (c *Calculator) RunCalculationsPopulateOutputFields() {
	c.copyInputIntoOutput()
	c.getTipoutPercentagesAndPoolInfo()
	c.tallyTipPools()
	c.distributeTipoutsGetFinalPayouts()
}

// helper methods
func (c *Calculator) copyInputIntoOutput() {
	// initialize output slices w/ same length as input slices
	c.BarTeamOut.Bartenders = make([]BartenderOut, len(c.BarTeamIn.Bartenders))
	c.ServersOut = make([]ServerOut, len(c.ServersIn))
	c.EventsOut = make([]EventOut, len(c.EventsIn))
	c.SupportOut = make([]SupportOut, len(c.SupportIn))

	// bar
	c.BarTeamOut.OwedToPreTipout = c.BarTeamIn.OwedTo
	c.BarTeamOut.Sales = c.BarTeamIn.Sales

	if c.BarTeamIn.Bartenders != nil {
		for i, bartender := range c.BarTeamIn.Bartenders {
			c.BarTeamOut.Bartenders[i].Name = bartender.Name
			c.BarTeamOut.Bartenders[i].Hours = bartender.Hours
		}
	}

	// servers
	if c.ServersIn != nil {
		for i, server := range c.ServersIn {
			c.ServersOut[i].Name = server.Name
			c.ServersOut[i].Sales = server.Sales
			c.ServersOut[i].OwedToPreTipout = server.OwedTo
		}
	}

	// events
	if c.EventsIn != nil {
		for i, event := range c.EventsIn {
			c.EventsOut[i].Name = event.Name
			c.EventsOut[i].OwedToPreTipout = event.OwedTo
			c.EventsOut[i].Sales = event.Sales
			c.EventsOut[i].SplitBy = event.SplitBy
		}
	}

	// support
	if c.SupportIn != nil {
		for i, support := range c.SupportIn {
			c.SupportOut[i].Name = support.Name
			c.SupportOut[i].Hours = support.Hours
		}
	}
}

func (c *Calculator) getTipoutPercentagesAndPoolInfo() {
	// determine tipout percentages and counts for bar and support
	c.BarTipoutPercentage, c.BarCount = c.getTipoutPercentageToBar()
	c.SupportTipoutPercentage, c.SupportCount = c.getTipoutPercentageToSupport()

	// get total hours for bartenders/support
	c.TotalBarHours = c.getTotalBarHours()
	c.TotalSupportHours = c.getTotalSupportHours()
}

func (c *Calculator) tallyTipPools() {
	// bar pool
	// from servers
	for i := range c.ServersOut {
		server := &c.ServersOut[i]
		server.TipoutToBar = server.Sales * c.BarTipoutPercentage
		c.BarPool += server.TipoutToBar
		c.BarTeamOut.TipoutFromServers += server.TipoutToBar
	}
	// from events
	for _, event := range c.ServersOut {
		event.TipoutToBar = event.Sales * c.BarTipoutPercentage
		c.BarPool += event.TipoutToBar
		c.BarTeamOut.TipoutFromEvents += event.TipoutToBar
	}

	// support pool
	if c.SupportOut != nil {
		// from bar
		// calculate tipout to support and record in field
		c.BarTeamOut.TipoutToSupport = c.BarTeamOut.Sales * c.SupportTipoutPercentage
		// add it to the support pool running tally
		c.SupportPool += c.BarTeamOut.TipoutToSupport
		// vv kind of a redundant field now, but could be useful should tipout rules change
		c.BarTeamOut.TotalAmountTippedOut = c.BarTeamOut.TipoutToSupport
		c.ToSupportFromBarAmount += c.BarTeamOut.TipoutToSupport

		// from servers
		for i := range c.ServersOut {
			server := &c.ServersOut[i]
			server.TipoutToSupport = server.Sales * c.SupportTipoutPercentage
			c.SupportPool += server.TipoutToSupport
			c.ToSupportFromServersAmount += server.TipoutToSupport
		}
		// from events
		for i := range c.EventsOut {
			event := &c.EventsOut[i]
			event.TipoutToSupport = event.Sales * c.SupportTipoutPercentage
			c.SupportPool += event.TipoutToSupport
			c.ToSupportFromEventsAmount += event.TipoutToSupport
		}
	}
}

func (c *Calculator) distributeTipoutsGetFinalPayouts() {
	// bar team
	c.BarTeamOut.FinalPayout = c.BarTeamOut.OwedToPreTipout - c.BarTeamOut.TotalAmountTippedOut + c.BarPool
	// bartenders
	if c.BarTeamOut.Bartenders != nil {
		// for _, bartender := range c.BarTeamOut.Bartenders {
		for i := range c.BarTeamOut.Bartenders {
			bartender := &c.BarTeamOut.Bartenders[i]
			bartender.PercentageOfBarTipPool = bartender.Hours / c.TotalBarHours
			bartender.OwedToPreTipout = c.BarTeamOut.OwedToPreTipout * bartender.PercentageOfBarTipPool
			bartender.TipoutToSupport = c.BarTeamOut.TipoutToSupport * bartender.PercentageOfBarTipPool
			bartender.TotalAmountTippedOut = bartender.TipoutToSupport
			bartender.TipoutFromServers = c.BarTeamOut.TipoutFromServers * bartender.PercentageOfBarTipPool
			bartender.TipoutFromEvents = c.BarTeamOut.TipoutFromEvents * bartender.PercentageOfBarTipPool
			bartender.TotalTipoutReceived = bartender.TipoutFromServers + bartender.TipoutFromEvents
			bartender.FinalPayout = bartender.OwedToPreTipout - bartender.TotalAmountTippedOut + (c.BarPool * bartender.PercentageOfBarTipPool)
		}
	}
	// servers
	for _, server := range c.ServersOut {
		server.TotalAmountTippedOut = server.TipoutToBar + server.TipoutToSupport
		server.FinalPayout = server.OwedToPreTipout - server.TotalAmountTippedOut
	}
	// events
	for _, event := range c.EventsOut {
		event.TotalAmountTippedOut = event.TipoutToBar + event.TipoutToSupport
		event.FinalPayout = event.OwedToPreTipout - event.TotalAmountTippedOut
	}
	// support
	for _, support := range c.SupportOut {
		support.FinalPayout = c.SupportPool * support.PercentageOfSupportTipPool
	}
}

func (c *Calculator) getTotalBarHours() float64 {
	totalHours := 0.0
	for _, bartender := range c.BarTeamIn.Bartenders {
		totalHours += bartender.Hours
	}
	return totalHours
}

func (c *Calculator) getTotalSupportHours() float64 {
	totalHours := 0.0
	for _, support := range c.SupportIn {
		totalHours += support.Hours
	}
	return totalHours
}

func (c *Calculator) getTipoutPercentageToBar() (float64, int) {
	count := len(c.BarTeamIn.Bartenders)
	var percentage float64
	if count == 0 {
		percentage = 0.00
	} else if count >= 3 {
		percentage = 0.015
	} else {
		percentage = 0.02
	}
	return percentage, count
}

func (c *Calculator) getTipoutPercentageToSupport() (float64, int) {
	count := len(c.SupportIn)
	var percentage float64
	if count == 0 {
		percentage = 0.00
	}
	if count <= 3 {
		percentage = float64(count) * 0.01
	} else {
		percentage = 0.03
	}
	return percentage, count
}
