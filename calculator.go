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

	// fields needed for calculations
	BarPool             float64
	TotalBarHours       float64
	BarCount            int
	BarTipoutPercentage float64

	SupportPool             float64
	TotalSupportHours       float64
	SupportCount            int
	SupportTipoutPercentage float64
}

func (c *Calculator) RunCalculationsPopulateOutputFields() {
	c.copyInputIntoOutput()
	c.getTipoutPercentages()
	c.tallyTipPools()
	c.distributeTipoutsGetFinalPayouts()
}

// helper funcs
func (c *Calculator) getTipoutPercentages() {
	// determine tipout percentages and counts for bar and support
	c.BarTipoutPercentage, c.BarCount = c.getTipoutPercentageToBar()
	c.SupportTipoutPercentage, c.SupportCount = c.getTipoutPercentageToSupport()

	// get total hours for bartenders/support
	c.TotalBarHours = c.getTotalBarHours()
	c.TotalSupportHours = c.getTotalSupportHours()

}

func (c *Calculator) distributeTipoutsGetFinalPayouts() {
	if c.BarTeamOut.Bartenders != nil {
		for _, bartender := range c.BarTeamOut.Bartenders {
			bartender.PercentageOfBarTipPool = bartender.Hours / c.TotalBarHours
			bartender.OwedToPreTipout = c.BarTeamOut.OwedToPreTipout * bartender.PercentageOfBarTipPool
			bartender.TipoutToSupport = c.BarTeamOut.TipoutToSupport * bartender.PercentageOfBarTipPool
			bartender.TotalAmountTippedOut = bartender.TipoutToSupport
		}
	}
}

func (c *Calculator) tallyTipPools() {
	// bar pool
	//servers
	for _, server := range c.ServersOut {
		server.TipoutToBar = server.Sales * c.BarTipoutPercentage
		c.BarPool += server.TipoutToBar
	}
	//events
	for _, event := range c.ServersOut {
		event.TipoutToBar = event.Sales * c.BarTipoutPercentage
		c.BarPool += event.TipoutToBar
	}

	// support pool
	if c.SupportOut != nil {
		//bar
		// calculate tipout to support and record in field
		c.BarTeamOut.TipoutToSupport = c.BarTeamOut.Sales * c.SupportTipoutPercentage
		// add it to the support pool running tally
		c.SupportPool += c.BarTeamOut.TipoutToSupport
		// vv kind of a redundant field now, but could be useful should tipout rules change
		c.BarTeamOut.TotalAmountTippedOut = c.BarTeamOut.TipoutToSupport

		//servers
		for _, server := range c.ServersOut {
			server.TipoutToSupport = server.Sales * c.SupportTipoutPercentage
			c.SupportPool += server.TipoutToSupport
		}
		//events
		for _, event := range c.EventsOut {
			event.TipoutToSupport = event.Sales * c.SupportTipoutPercentage
			c.SupportPool += event.TipoutToSupport
		}
	}
}

func (c *Calculator) copyInputIntoOutput() {
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
