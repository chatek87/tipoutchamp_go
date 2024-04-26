package main

type Calculator struct {
	// input fields
	BarTeamIn     BarTeamIn
	ServersIn     []ServerIn
	EventsIn      []EventIn
	SupportTeamIn []SupportIn

	// output fields
	BarTeamOut     BarTeamOut
	ServersOut     []ServerOut
	EventsOut      []EventOut
	SupportTeamOut []SupportOut
}

func (c *Calculator) GetTotalBarHours() float64 {
	totalHours := 0.0
	for _, bartender := range c.BarTeamIn.Bartenders {
		totalHours += bartender.Hours
	}
	return totalHours
}

func (c *Calculator) GetTotalSupportHours() float64 {
	totalHours := 0.0
	for _, support := range c.SupportTeamIn {
		totalHours += support.Hours
	}
	return totalHours
}

func (c *Calculator) GetTipoutPercentageToBar() float64 {
	supportCount := len(c.SupportTeamIn)
	if supportCount == 0 {
		return 0.00
	}
	if supportCount >= 3 {
		return 0.015
	} else {
		return 0.02
	}
}

func (c *Calculator) GetTipoutPercentageToSupport() float64 {
	supportCount := len(c.SupportTeamIn)
	if supportCount == 0 {
		return 0.00
	}
	if supportCount <= 3 {
		return float64(supportCount) * 0.01
	} else {
		return 0.03
	}
}
