package onoffbasic

import (
	"controllers/def/info"
	"fmt"
	"os"
	"shared"
)

type Controller struct {
	Info info.Controller
}

func (c *Controller) Initialise(p ...float64) {

	if len(p) < 2 {
		fmt.Printf("Error: '%s controller requires 2 info (min,max) \n", shared.BASIC_ONOFF)
		os.Exit(0)
	}

	c.Info.TypeName = shared.BASIC_ONOFF
	c.Info.Min = p[0]
	c.Info.Max = p[1]
}

func (c *Controller) Update(p ...float64) float64 {

	direction := -1.0
	u := 0.0

	s := p[0] // goal
	y := p[1] // plant output

	// error
	err := direction * (s - y)

	// control law
	if err >= 0 {
		u = c.Info.Max
	} else {
		u = c.Info.Min
	}
	return u
}
