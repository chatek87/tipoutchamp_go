package main

import "fmt"

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

	// configuration fields
	BarCount                int
	SupportCount            int
	TotalBarHours           float64
	TotalSupportHours       float64
	BarTipoutPercentage     float64
	SupportTipoutPercentage float64

	// tip pool-related fields
	BarPool     float64
	SupportPool float64

	ToSupportFromBarAmount     float64
	ToSupportFromServersAmount float64
	ToSupportFromEventsAmount  float64
}

func (c *Calculator) RunCalculationsPopulateOutputFields() {
	c.copyInputIntoOutput()
	c.setConfigurationFields()
	c.tallyTipPools()
	c.distributeTipoutsGetFinalPayouts()
}

// helpers to RunCalculationsPopulateOutputFields()
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
func (c *Calculator) setConfigurationFields() {
	// counts
	c.BarCount = len(c.BarTeamIn.Bartenders)
	c.SupportCount = len(c.SupportIn)
	// tipout %'s
	c.setBarTipoutPercentage()
	c.setSupportTipoutPercentage()
	// hours
	c.setTotalBarHours()
	c.setTotalSupportHours()
}
func (c *Calculator) setTotalBarHours() {
	totalHours := 0.0
	for _, bartender := range c.BarTeamIn.Bartenders {
		totalHours += bartender.Hours
	}
	c.TotalBarHours = totalHours
}
func (c *Calculator) setTotalSupportHours() {
	totalHours := 0.0
	for _, support := range c.SupportIn {
		totalHours += support.Hours
	}
	c.TotalSupportHours = totalHours
}
func (c *Calculator) setBarTipoutPercentage() {
	count := len(c.SupportIn)
	if count >= 3 {
		c.BarTipoutPercentage = 0.015
	} else {
		c.BarTipoutPercentage = 0.02
	}
}
func (c *Calculator) setSupportTipoutPercentage() {
	count := len(c.SupportIn)
	if count == 0 {
		c.SupportTipoutPercentage = 0.00
	}
	if count <= 3 {
		c.SupportTipoutPercentage = float64(count) * 0.01
	} else {
		c.SupportTipoutPercentage = 0.03
	}
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
	for i := range c.EventsOut {
		event := &c.EventsOut[i]
		event.TipoutToBar = event.Sales * c.BarTipoutPercentage
		c.BarPool += event.TipoutToBar
		c.BarTeamOut.TipoutFromEvents += event.TipoutToBar
	}

	c.BarTeamOut.TotalTipoutReceived = c.BarTeamOut.TipoutFromServers + c.BarTeamOut.TipoutFromEvents

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
	for i := range c.ServersOut {
		server := &c.ServersOut[i]
		server.TotalAmountTippedOut = server.TipoutToBar + server.TipoutToSupport
		server.FinalPayout = server.OwedToPreTipout - server.TotalAmountTippedOut
	}
	// events
	for i := range c.EventsOut {
		event := &c.EventsOut[i]
		event.TotalAmountTippedOut = event.TipoutToBar + event.TipoutToSupport
		event.FinalPayout = event.OwedToPreTipout - event.TotalAmountTippedOut
		event.FinalPayoutPerWorker = event.FinalPayout / float64(event.SplitBy)
	}
	// support
	for i := range c.SupportOut {
		support := &c.SupportOut[i]
		support.PercentageOfSupportTipPool = support.Hours / c.TotalSupportHours
		support.TipoutFromBar = c.ToSupportFromBarAmount * support.PercentageOfSupportTipPool
		support.TipoutFromServers = c.ToSupportFromServersAmount * support.PercentageOfSupportTipPool
		support.TipoutFromEvents = c.ToSupportFromEventsAmount * support.PercentageOfSupportTipPool
		support.FinalPayout = c.SupportPool * support.PercentageOfSupportTipPool
	}
}

// report related
func (c *Calculator) GenerateReport() {
	fmt.Printf("TIPOUT REPORT\n")
	fmt.Printf("\n")
	fmt.Printf("BAR TEAM\n")
	fmt.Printf("OWED TO PRE TIPOUT: %.2f\n", c.BarTeamOut.OwedToPreTipout)
	fmt.Printf("SALES: %.2f\n", c.BarTeamOut.Sales)
	fmt.Printf("TIPOUT TO SUPPORT: %.2f\n", c.BarTeamOut.TipoutToSupport)
	fmt.Printf("TOTAL AMOUNT TIPPED OUT: %.2f\n", c.BarTeamOut.TotalAmountTippedOut)
	fmt.Printf("TIPOUT FROM SERVERS: %.2f\n", c.BarTeamOut.TipoutFromServers)
	fmt.Printf("TIPOUT FROM EVENTS: %.2f\n", c.BarTeamOut.TipoutFromEvents)
	fmt.Printf("TOTAL TIPOUT RECEIVED: %.2f\n", c.BarTeamOut.TotalTipoutReceived)
	fmt.Printf("FINAL PAYOUT: %.2f\n", c.BarTeamOut.FinalPayout)
	fmt.Printf("INDIVIDUAL BARTENDERS\n")
	fmt.Printf("\n")
	for _, b := range c.BarTeamOut.Bartenders {
		fmt.Printf("BARTENDER NAME: %s\n", b.Name)
		fmt.Printf("HOURS: %.2f\n", b.Hours)
		fmt.Printf("%% OF BAR TIP POOL: %.2f\n", b.PercentageOfBarTipPool)
		fmt.Printf("OWED TO PRE TIPOUT: %.2f\n", b.OwedToPreTipout)
		fmt.Printf("TIPOUT TO SUPPORT: %.2f\n", b.TipoutToSupport)
		fmt.Printf("TOTAL AMOUNT TIPPED OUT: %.2f\n", b.TotalAmountTippedOut)
		fmt.Printf("TIPOUT FROM SERVERS: %.2f\n", b.TipoutFromServers)
		fmt.Printf("TIPOUT FROM EVENTS: %.2f\n", b.TipoutFromEvents)
		fmt.Printf("TOTAL TIPOUT RECEIVED: %.2f\n", b.TotalTipoutReceived)
		fmt.Printf("FINAL PAYOUT: %.2f\n", b.FinalPayout)
		fmt.Printf("\n")
	}
	fmt.Printf("SERVERS\n")
	for _, s := range c.ServersOut {
		fmt.Printf("SERVER NAME: %s\n", s.Name)
		fmt.Printf("OWED TO PRE TIPOUT: %.2f\n", s.OwedToPreTipout)
		fmt.Printf("SALES: %.2f\n", s.Sales)
		fmt.Printf("TIPOUT TO BAR: %.2f\n", s.TipoutToBar)
		fmt.Printf("TIPOUT TO SUPPORT: %.2f\n", s.TipoutToSupport)
		fmt.Printf("TOTAL AMOUNT TIPPED OUT: %.2f\n", s.TotalAmountTippedOut)
		fmt.Printf("FINAL PAYOUT: %.2f\n", s.FinalPayout)
		fmt.Printf("\n")
	}
	fmt.Printf("EVENTS\n")
	for _, e := range c.EventsOut {
		fmt.Printf("EVENT NAME: %s\n", e.Name)
		fmt.Printf("OWED TO PRE TIPOUT: %.2f\n", e.OwedToPreTipout)
		fmt.Printf("SALES: %.2f\n", e.Sales)
		fmt.Printf("SPLIT BY: %d\n", e.SplitBy)
		fmt.Printf("TIPOUT TO BAR: %.2f\n", e.TipoutToBar)
		fmt.Printf("TIPOUT TO SUPPORT: %.2f\n", e.TipoutToSupport)
		fmt.Printf("TOTAL AMOUNT TIPPED OUT: %.2f\n", e.TotalAmountTippedOut)
		fmt.Printf("FINAL PAYOUT: %.2f\n", e.FinalPayout)
		fmt.Printf("FINAL PAYOUT PER WORKER: %.2f\n", e.FinalPayoutPerWorker)
		fmt.Printf("\n")
	}
	fmt.Printf("SUPPORT\n")
	for _, s := range c.SupportOut {
		fmt.Printf("SUPPORT NAME: %s\n", s.Name)
		fmt.Printf("HOURS: %.2f\n", s.Hours)
		fmt.Printf("%% OF SUPPORT TIP POOL: %.2f\n", s.PercentageOfSupportTipPool)
		fmt.Printf("TIPOUT FROM BAR: %.2f\n", s.TipoutFromBar)
		fmt.Printf("TIPOUT FROM SERVERS: %.2f\n", s.TipoutFromServers)
		fmt.Printf("TIPOUT FROM EVENTS: %.2f\n", s.TipoutFromEvents)
		fmt.Printf("FINAL PAYOUT: %.2f\n", s.FinalPayout)
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (c *Calculator) seedSampleData() {
	// Define sample data for employees
	c.BarTeamIn = BarTeamIn{Bartenders: []BartenderIn{{Name: "Bartender 1", Hours: 6.0}, {Name: "Bartender 2", Hours: 8.0}}, OwedTo: 400.00, Sales: 2000.00}
	c.ServersIn = []ServerIn{{Name: "Server 1", OwedTo: 100.00, Sales: 500.00}, {Name: "Server 2", OwedTo: 200.00, Sales: 1000.00}, {Name: "Server 3", OwedTo: 300.00, Sales: 1500.00}, {Name: "Server 4", OwedTo: 400.00, Sales: 2000.00}}
	c.EventsIn = []EventIn{{Name: "Event 1", OwedTo: 600.00, Sales: 3000.00, SplitBy: 2}}
	c.SupportIn = []SupportIn{{Name: "Support 1", Hours: 4.0}, {Name: "Support 2", Hours: 5.0}, {Name: "Support 3", Hours: 6.0}}
}

func getSampleCalc() Calculator {
	calc := Calculator{
		BarTeamIn: BarTeamIn{Bartenders: []BartenderIn{{Name: "Bartender 1", Hours: 6.0}, {Name: "Bartender 2", Hours: 8.0}}, OwedTo: 400.00, Sales: 2000.00},
		ServersIn: []ServerIn{{Name: "Server 1", OwedTo: 100.00, Sales: 500.00}, {Name: "Server 2", OwedTo: 200.00, Sales: 1000.00}, {Name: "Server 3", OwedTo: 300.00, Sales: 1500.00}, {Name: "Server 4", OwedTo: 400.00, Sales: 2000.00}},
		EventsIn:  []EventIn{{Name: "Event 1", OwedTo: 600.00, Sales: 3000.00, SplitBy: 2}},
		SupportIn: []SupportIn{{Name: "Support 1", Hours: 4.0}, {Name: "Support 2", Hours: 5.0}, {Name: "Support 3", Hours: 6.0}},
	}
	return calc
}

// func (c *Calculator) SaveJSONToFile() {

// }
