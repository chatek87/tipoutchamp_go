package main

import (
	"testing"
)

func TestCopyInputIntoOutput(t *testing.T) {
	c := testSetup()

	c.copyInputIntoOutput()
	// bar
	if c.BarTeamOut.OwedToPreTipout != 400 || c.BarTeamIn.OwedTo != 400 {
		t.Errorf("Expected BarTeamIn.OwedTo == BarTeamOut.OwedToPreTipout.\nGot: (in) %.02f, (out) %.02f", c.BarTeamIn.OwedTo, c.BarTeamOut.OwedToPreTipout)
	}
	if c.BarTeamOut.Sales != 2000.00 || c.BarTeamIn.Sales != 2000.00 {
		t.Errorf("Expected BarTeamIn.OwedTo == BarTeamOut.OwedToPreTipout.\nGot: (in) %.02f, (out) %.02f", c.BarTeamIn.OwedTo, c.BarTeamOut.OwedToPreTipout)
	}
	// bartenders
	if len(c.BarTeamOut.Bartenders) != len(c.BarTeamIn.Bartenders) {
		t.Errorf("Expected number of Bartenders to be the same.\nGot: (in) %d, (out) %d", len(c.BarTeamIn.Bartenders), len(c.BarTeamOut.Bartenders))
	}

	for i, bartenderIn := range c.BarTeamIn.Bartenders {
		bartenderOut := c.BarTeamOut.Bartenders[i]
		if bartenderIn.Name != bartenderOut.Name || bartenderIn.Hours != bartenderOut.Hours {
			t.Errorf("Expected BartenderIn and BartenderOut to be the same at index %d.\nGot: (in) %+v, (out) %+v", i, bartenderIn, bartenderOut)
		}
	}
	// servers
	if len(c.ServersIn) != len(c.ServersOut) {
		t.Errorf("Expected number of Servers to be the same.\nGot: (in) %d, (out) %d", len(c.ServersIn), len(c.ServersOut))
	}

	for i, serverIn := range c.ServersIn {
		serverOut := c.ServersOut[i]
		if serverIn.Name != serverOut.Name || serverIn.OwedTo != serverOut.OwedToPreTipout || serverIn.Sales != serverOut.Sales {
			t.Errorf("Expected ServerIn and ServerOut to be the same at index %d.\nGot: (in) %+v, (out) %+v", i, serverIn, serverOut)
		}
	}
	// events
	if len(c.EventsIn) != len(c.EventsOut) {
		t.Errorf("Expected number of Events to be the same.\nGot: (in) %d, (out) %d", len(c.EventsIn), len(c.EventsOut))
	}

	for i, eventIn := range c.EventsIn {
		eventOut := c.EventsOut[i]
		if eventIn.Name != eventOut.Name || eventIn.OwedTo != eventOut.OwedToPreTipout || eventIn.Sales != eventOut.Sales || eventIn.SplitBy != eventOut.SplitBy {
			t.Errorf("Expected EventIn and EventOut to be the same at index %d.\nGot: (in) %+v, (out) %+v", i, eventIn, eventOut)
		}
	}
	// support
	if len(c.SupportIn) != len(c.SupportOut) {
		t.Errorf("Expected number of Support to be the same.\nGot: (in) %d, (out) %d", len(c.SupportIn), len(c.SupportOut))
	}

	for i, supportIn := range c.SupportIn {
		supportOut := c.SupportOut[i]
		if supportIn.Hours != supportOut.Hours || supportIn.Name != supportOut.Name {
			t.Errorf("Expected SupportIn and SupportOut to be the same at index %d.\nGot: (in) %+v, (out) %+v", i, supportIn, supportOut)
		}
	}
}

func TestGetTipoutPercentageToBarAndToSupport(t *testing.T) {
	c0 := testSetup() // 0 support
	c0.SupportIn = []SupportIn{}

	c1 := testSetup() // 1 support
	c1.SupportIn = []SupportIn{{Name: "Support 1", Hours: 4.0}}

	c2 := testSetup() // 2 support
	c2.SupportIn = []SupportIn{{Name: "Support 1", Hours: 4.0}, {Name: "Support 2", Hours: 5.0}}

	c3 := testSetup() // 3 support
	c3.SupportIn = []SupportIn{{Name: "Support 1", Hours: 4.0}, {Name: "Support 2", Hours: 5.0}, {Name: "Support 3", Hours: 6.0}}

	c4 := testSetup() // more than 3 support
	newSupport := SupportIn{Name: "Support 4", Hours: 3.0}
	c4.SupportIn = append(c4.SupportIn, newSupport)

	calcs := []Calculator{c0, c1, c2, c3, c4}
	for _, c := range calcs {
		// setup necessary fields
		c.copyInputIntoOutput()
		c.setConfigurationFields()
		// check bar tipout percentage
		if c.SupportCount >= 3 && c.BarTipoutPercentage != 0.015 {
			t.Errorf("BarTipoutPercentage expected: 0.015 got: %f", c.BarTipoutPercentage)
		} else if c.SupportCount < 3 && c.BarTipoutPercentage != 0.02 {
			t.Errorf("BarTipoutPercentage expected: 0.015 got: %f", c.BarTipoutPercentage)
		}
		// check support tipout percentage
		if c.SupportCount == 0 && c.SupportTipoutPercentage != 0.00 {
			t.Errorf("SupportTipoutPercentage expected: 0.00 got: %f", c.SupportTipoutPercentage)
		} else if c.SupportCount == 1 && c.SupportTipoutPercentage != 0.01 {
			t.Errorf("SupportTipoutPercentage expected: 0.01 got: %f", c.SupportTipoutPercentage)
		} else if c.SupportCount == 2 && c.SupportTipoutPercentage != 0.02 {
			t.Errorf("SupportTipoutPercentage expected: 0.02 got: %f", c.SupportTipoutPercentage)
		} else if c.SupportCount >= 3 && c.SupportTipoutPercentage != 0.03 {
			t.Errorf("SupportTipoutPercentage expected: 0.03 got: %f", c.SupportTipoutPercentage)
		}
	}
}

// helper funcs
func testSetup() Calculator {
	calc := Calculator{
		BarTeamIn: BarTeamIn{Bartenders: []BartenderIn{{Name: "Bartender 1", Hours: 6.0}, {Name: "Bartender 2", Hours: 8.0}}, OwedTo: 400.00, Sales: 2000.00},
		ServersIn: []ServerIn{{Name: "Server 1", OwedTo: 100.00, Sales: 500.00}, {Name: "Server 2", OwedTo: 200.00, Sales: 1000.00}, {Name: "Server 3", OwedTo: 300.00, Sales: 1500.00}, {Name: "Server 4", OwedTo: 400.00, Sales: 2000.00}},
		EventsIn:  []EventIn{{Name: "Event 1", OwedTo: 600.00, Sales: 3000.00, SplitBy: 2}},
		SupportIn: []SupportIn{{Name: "Support 1", Hours: 4.0}, {Name: "Support 2", Hours: 5.0}, {Name: "Support 3", Hours: 6.0}},
	}
	return calc
}
