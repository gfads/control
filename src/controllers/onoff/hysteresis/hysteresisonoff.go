package hysteresisonoff

import (
	"fmt"
	"os"
	"shared"
)

type Controller struct {
	Min            float64
	Max            float64
	HysteresisBand float64
	PreviousOut    float64
}

func (c *Controller) Initialise(p ...float64) {

	if len(p) < 3 {
		fmt.Printf("Error: '%s' controller requires 3 parameters (min,max,hysteresis band) \n", shared.HYSTERESIS_ONOFF)
		os.Exit(0)
	}
	c.Min = p[0]
	c.Max = p[1]
	c.HysteresisBand = p[2]
	c.PreviousOut = 0.0
}

func (c *Controller) Update(p ...float64) float64 {

	direction := -1.0
	u := 0.0

	s := p[0] // goal
	y := p[1] // plant output

	// error
	err := direction * (s - y)

	// control law
	if err > -c.HysteresisBand/2.0 && err < c.HysteresisBand/2.0 {
		u = c.PreviousOut
	}
	if err >= c.HysteresisBand/2.0 {
		u = c.Max
	}
	if err <= -c.HysteresisBand/2.0 {
		u = c.Min
	}

	if u < c.Min {
		u = c.Min
	}
	if u > c.Max {
		u = c.Max
	}
	c.PreviousOut = u

	return u
}
