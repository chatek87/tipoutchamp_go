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
}

func (c *Calculator) RunCalculationsPopulateOutputFields() {
	// copy all input fields into output model fields
	c.CopyInputIntoOutput()

	// determine tipout percentages and counts for bar and support
	barTipoutPercentage, barCount := c.getTipoutPercentageToBar()
	supportTipoutPercentage, supportCount := c.getTipoutPercentageToSupport()

	// get total hours for bartenders/support
	barHours := c.getTotalBarHours()
	supportHours := c.getTotalSupportHours()

	// go thru each server, calculating tipout to bar and support, recording that value in the corresponding Out fields
	// subtract the calculated tipout from the OwedTo field to get the FinalPayout, record that value
	// add the calculated tipout to a running tally of total tip pool for
}

// helper funcs
func (c *Calculator) CopyInputIntoOutput() {
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
