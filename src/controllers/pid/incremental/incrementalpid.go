package incrementalpid

import (
	"controllers/info"
	"fmt"
	"os"
	"shared"
)

const DeltaTime = 1 // see page 103

type Controller struct {
	Info info.InfoController
}

func (c *Controller) Initialise(p ...float64) {

	if len(p) < 5 {
		fmt.Printf("Error: '%s' controller requires 5 parameters (min,max,kp,ki,kd) \n", shared.INCREMENTAL_FORM_PID)
		os.Exit(0)
	}

	c.Info.Min = p[0]
	c.Info.Max = p[1]

	c.Info.Kp = p[2]
	c.Info.Ki = p[3]
	c.Info.Kd = p[4]

	c.Info.Integrator = 0.0
	c.Info.PreviousError = 0.0
	c.Info.PreviousPreviousError = 0.0
	c.Info.SumPreviousErrors = 0.0
	c.Info.Out = 0.0
	c.Info.PreviousDifferentiator = 0.0
}

func (c *Controller) Update(p ...float64) float64 {
	r := p[0] // goal
	y := p[1] // plant output

	// errors
	err := r - y

	// Integrator // page 106
	c.Info.Integrator += DeltaTime * err
	integrator := c.Info.Integrator * c.Info.Ki

	// Delta of the new PC
	deltaU := c.Info.Kp*(err-c.Info.PreviousError) + c.Info.Ki*err*DeltaTime + c.Info.Kd*(err-2*c.Info.PreviousError+c.Info.PreviousPreviousError)/DeltaTime

	// pid output
	c.Info.Out = integrator + deltaU // see page 106 why add an integrator

	if c.Info.Out > c.Info.Max {
		c.Info.Out = c.Info.Max
	} else if c.Info.Out < c.Info.Min {
		c.Info.Out = c.Info.Min
	}

	c.Info.PreviousPreviousError = c.Info.PreviousError
	c.Info.PreviousError = err
	c.Info.SumPreviousErrors = c.Info.SumPreviousErrors + err

	return c.Info.Out
}