package main

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

	// Optionally, call methods on the Calculator instance
	// For example, to run calculations and populate output fields
	// calc.RunCalculationsPopulateOutputFields()
}
