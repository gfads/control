/*********************************************************************************
Author: Nelson S Rosa
Description: This program implements the On Off controller with hysteresis as defined
			in "Feedback Control for Computer Systems: Introducing Control Theory to
			Enterprise Programmers", Philipp K. Janert, 2014.
Date: 04/02/2023
*********************************************************************************/

package hysteresisonoff

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

	if len(p) < 3 {
		fmt.Printf("Error: '%s' controller requires 3 info (min,max,hysteresis band) \n", shared.HYSTERESIS_ONOFF)
		os.Exit(0)
	}
	c.Info.Min = p[0]
	c.Info.Max = p[1]
	c.Info.HysteresisBand = p[2]
	c.Info.PreviousOut = 0.0
}

func (c *Controller) Update(p ...float64) float64 {

	direction := -1.0
	u := 0.0

	s := p[0] // goal
	y := p[1] // plant output

	// error
	err := direction * (s - y)

	// control law
	if err > -c.Info.HysteresisBand/2.0 && err < c.Info.HysteresisBand/2.0 {
		u = c.Info.PreviousOut
	}
	if err >= c.Info.HysteresisBand/2.0 {
		u = c.Info.Max
	}
	if err <= -c.Info.HysteresisBand/2.0 {
		u = c.Info.Min
	}

	if u < c.Info.Min {
		u = c.Info.Min
	}
	if u > c.Info.Max {
		u = c.Info.Max
	}
	c.Info.PreviousOut = u

	return u
}
