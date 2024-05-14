package main

import (
	"testing"
)

func TestCopyInputIntoOutput(t *testing.T) {
	calc := Calculator{
		BarTeamIn: BarTeamIn{Bartenders: []BartenderIn{{Name: "Bartender 1", Hours: 6.0}, {Name: "Bartender 2", Hours: 8.0}}, OwedTo: 400.00, Sales: 2000.00},
		ServersIn: []ServerIn{{Name: "Server 1", OwedTo: 100.00, Sales: 500.00}, {Name: "Server 2", OwedTo: 200.00, Sales: 1000.00}, {Name: "Server 3", OwedTo: 300.00, Sales: 1500.00}, {Name: "Server 4", OwedTo: 400.00, Sales: 2000.00}},
		EventsIn:  []EventIn{{Name: "Event 1", OwedTo: 600.00, Sales: 3000.00, SplitBy: 2}},
		SupportIn: []SupportIn{{Name: "Support 1", Hours: 4.0}, {Name: "Support 2", Hours: 5.0}, {Name: "Support 3", Hours: 6.0}},
	}

	calc.copyInputIntoOutput()
	// bar
	if calc.BarTeamOut.OwedToPreTipout != 400 || calc.BarTeamIn.OwedTo != 400 {
		t.Errorf("Expected BarTeamIn.OwedTo == BarTeamOut.OwedToPreTipout.\nGot: (in) %.02f, (out) %.02f", calc.BarTeamIn.OwedTo, calc.BarTeamOut.OwedToPreTipout)
	}
	if calc.BarTeamOut.Sales != 2000.00 || calc.BarTeamIn.Sales != 2000.00 {
		t.Errorf("Expected BarTeamIn.OwedTo == BarTeamOut.OwedToPreTipout.\nGot: (in) %.02f, (out) %.02f", calc.BarTeamIn.OwedTo, calc.BarTeamOut.OwedToPreTipout)
	}
	// Check Bartenders
	if len(calc.BarTeamOut.Bartenders) != len(calc.BarTeamIn.Bartenders) {
		t.Errorf("Expected number of Bartenders to be the same.\nGot: (in) %d, (out) %d", len(calc.BarTeamIn.Bartenders), len(calc.BarTeamOut.Bartenders))
	}

	for i, bartenderIn := range calc.BarTeamIn.Bartenders {
		bartenderOut := calc.BarTeamOut.Bartenders[i]
		if bartenderIn.Name != bartenderOut.Name || bartenderIn.Hours != bartenderOut.Hours {
			t.Errorf("Expected BartenderIn and BartenderOut to be the same at index %d.\nGot: (in) %+v, (out) %+v", i, bartenderIn, bartenderOut)
		}
	}
	// servers
	for i, serverIn := range calc.ServersIn {
		serverOut := calc.ServersOut[i]
		if serverIn.Name != serverOut.Name {
			t.Errorf("Expected ServerIn.Name and ServerOut.Name to be the same.")
		}
	}
	// events
	// support
}
