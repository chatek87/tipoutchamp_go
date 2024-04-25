package main

type Calculator struct {
	BarTeam BarTeam
	Servers []Server
	Events  []Event
	SupportTeam []Support

	TotalBarHours float64
	TotalSupportHours float64
}

func (c *Calculator) SetTotalBarHours() {
	totalHours := 0.0
	for _, bartender := range c.BarTeam.Bartenders {
		totalHours += bartender.Hours
	}
	c.TotalBarHours = totalHours
}

func (c *Calculator) SetTotalSupportHours() {
	totalHours := 0.0
	for _, support := range c.SupportTeam {
		totalHours += support.Hours
	}
}