package main

import "fmt"

func main() {
	seedSampleData()
}

func seedSampleData() {
	// Define sample data for employees
	barTeamIn := BarTeamIn{
		// Populate with sample data
		Bartenders: []BartenderIn{
			{
				Name:  "Bartender 1",
				Hours: 6.0,
			},
			{
				Name:  "Bartender 2",
				Hours: 8.0,
			},
		},
		OwedTo: 400.00,
		Sales:  2000.00,
	}

	serversIn := []ServerIn{
		// Populate with sample data
		{
			Name:   "Server 1",
			OwedTo: 100.00,
			Sales:  500.00,
		},
		{
			Name:   "Server 2",
			OwedTo: 200.00,
			Sales:  1000.00,
		},
		{
			Name:   "Server 3",
			OwedTo: 300.00,
			Sales:  1500.00,
		},
		{
			Name:   "Server 4",
			OwedTo: 400.00,
			Sales:  2000.00,
		},
	}

	eventsIn := []EventIn{
		// Populate with sample data
		{
			Name:    "Event 1",
			OwedTo:  600.00,
			Sales:   3000.00,
			SplitBy: 2,
		},
	}

	supportIn := []SupportIn{
		// Populate with sample data
		{
			Name:  "Support 1",
			Hours: 4.0,
		},
		{
			Name:  "Support 2",
			Hours: 5.0,
		},
		{
			Name:  "Support 3",
			Hours: 6.0,
		},
	}

	// Initialize Calculator with sample data
	calc := Calculator{
		BarTeamIn: barTeamIn,
		ServersIn: serversIn,
		EventsIn:  eventsIn,
		SupportIn: supportIn,
	}

	calc.RunCalculationsPopulateOutputFields()

	printReport(calc)
}

func printReport(c Calculator) {
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
