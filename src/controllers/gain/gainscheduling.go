package gainscheduling

import (
	"controllers/info"
	"shared"
)

const DeltaTime = 1 // see page 103

type Controller struct {
	Info      info.InfoController
	GainTable [2][3]float64
}

func (c *Controller) Initialise(p ...float64) {

	// hard coded gain scheduling table
	c.GainTable[0][0] = -9600 // kp[0] // P
	c.GainTable[0][1] = 0.0   // ki[1]
	c.GainTable[0][2] = 0.0   // kd[2]

	//c.GainTable[1][0] = -9600 // kp[0] // PID
	//c.GainTable[1][1] = 0.5   // ki[1]
	//c.GainTable[1][2] = 0.01  // kd[2]

	c.GainTable[1][0] = -9600 // kp[0] // PI
	c.GainTable[1][1] = 0.5   // ki[1]
	c.GainTable[1][2] = 0.0   // kd[2]

	kp := c.GainTable[0][0]
	ki := c.GainTable[0][1]
	kd := c.GainTable[0][2]

	c.Info.Min = p[0]
	c.Info.Max = p[1]

	c.Info.Kp = kp
	c.Info.Ki = ki
	c.Info.Kd = kd

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

	// decide about the gain -- based on the capacitor energy level
	if y < shared.OV { // gain scheduling 1 (P Controller)
		c.Info.Kp = c.GainTable[0][0]
		c.Info.Ki = c.GainTable[0][1]
		c.Info.Kd = c.GainTable[0][2]
	} else { // gain scheduling 2 (PI Controller)
		c.Info.Kp = c.GainTable[1][0]
		c.Info.Ki = c.GainTable[1][1]
		c.Info.Kd = c.GainTable[1][2]
	}

	// Proportional
	proportional := c.Info.Kp * err

	// Integrator (David page 49)
	c.Info.Integrator += DeltaTime * err
	integrator := c.Info.Integrator * c.Info.Ki

	// Differentiator (David page 49)
	differentiator := c.Info.Kd * (err - c.Info.PreviousError) / DeltaTime

	// pid output
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