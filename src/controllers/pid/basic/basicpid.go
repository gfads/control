package basicpid

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
		fmt.Printf("Error: '%s' controller requires 5 parameters (min,max,kp,ki,kd) \n", shared.BASIC_PID)
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

	// Proportional
	proportional := c.Info.Kp * err

	// Integrator (page 49)
	c.Info.Integrator += DeltaTime * err
	integrator := c.Info.Integrator * c.Info.Ki

	// Differentiator (page 49)
	differentiator := c.Info.Kd * (err - c.Info.PreviousError) / DeltaTime

	// control law
	c.Info.Out = proportional + integrator + differentiator

	if c.Info.Out > c.Info.Max {
		c.Info.Out = c.Info.Max
	} else if c.Info.Out < c.Info.Min {
		c.Info.Out = c.Info.Min
	}

	c.Info.PreviousError = err
	c.Info.SumPreviousErrors = c.Info.SumPreviousErrors + err

	return c.Info.Out
}
