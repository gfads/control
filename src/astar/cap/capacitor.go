package cap

import (
	"fmt"
	"math/rand"
	"os"
	"shared"
)

type Capacitor struct {
	PreviousVoltageLevel float64
	Pattern              int
	PreviousPattern      int
	CountCycles          int
	State                int
	Indx                 int
}

var VoltageFromFile = []float64{5.156, 4.95,
	4.85,
	4.788,
	4.75,
	4.743,
	4.692,
	4.634,
	4.634,
	4.582,
	4.556,
	4.537,
	4.537,
	4.537,
	4.692,
	4.95,
	5.143,
	5.336,
	5.568,
	5.568,
	5.588,
	5.568,
	5.517,
	5.459,
	5.407,
	5.356,
	5.356,
	5.336,
	5.31,
	5.259,
	5.2,
	5.149,
	5.149,
	5.13,
	5.104,
	5.046,
	5.04,
	4.995,
	4.956,
	4.95,
	4.95,
	4.95,
	4.95,
	4.95,
	4.95,
	5.149,
	5.407,
	5.459,
	5.452,
	5.407,
	5.542,
	5.459,
	5.459,
	5.42,
	5.39,
	5.35,
	5.31,
	5.31,
	5.259,
	5.201,
	5.15,
	5.135,
	5.13,
	5.12,
	5.12,
	5.11,
	5.11,
	4.98,
	4.956,
	4.95,
	4.95,
	4.95,
	4.95,
	4.95,
	4.995,
	5,
	5.356,
	5.459,
	5.459,
	5.54,
	5.5,
	5.49,
	5.49,
	5.47,
	5.47,
	5.356,
	5.356,
	5.349,
	5.31,
	5.31,
	5.259,
	5.233,
	5.201,
	5.149,
	5.149,
	5.13,
	5.104,
	5.046,
	4.995,
	4.995,
	4.956,
	4.95,
	4.956,
	4.995,
	4.995,
	5.046,
	5.149,
	5.259,
	5.31,
	5.31,
	5.259,
	5.201,
	5.201,
	5.149,
	5.149,
	5.13,
	5.104,
	5.046,
	4.995,
	4.995,
	4.956,
	4.95,
	4.95,
	4.95,
	4.95,
	4.95,
	4.95,
	4.763,
	4.743,
	4.743,
	4.692,
	4.692,
}

func (c *Capacitor) GetCapacitorVoltageLevelFromFile() float64 {

	r := VoltageFromFile[c.Indx]
	c.Indx++

	if c.Indx >= len(VoltageFromFile) {
		c.Indx = 0
	}
	return r
}

func (c *Capacitor) GetCapacitorVoltageLevel() float64 {
	capacitorVoltage := 0.0
	pattern := c.Pattern

	if c.Pattern == shared.RandomHarvesting {
		switch c.State {
		case 0:
			pattern = rand.Intn(shared.RandomHarvesting)
			c.PreviousPattern = pattern
			c.State = 1
		case 1:
			if c.CountCycles < shared.SamplingCycleSize {
				pattern = c.PreviousPattern
				c.CountCycles++
			} else {
				c.CountCycles = 0
				for { // define a new pattern different from the previous one
					pattern = rand.Intn(shared.RandomHarvesting)
					if pattern != c.PreviousPattern {
						c.PreviousPattern = pattern
						break
					}
				}
			}
		}
	}

	switch pattern {
	case shared.IncreasingHarvesting:
		gain := shared.StepVoltage * float64(shared.RandInt(1, 100))
		if (c.PreviousVoltageLevel + gain) > shared.MaximumVoltage {
			capacitorVoltage = shared.MaximumVoltage
		} else {
			capacitorVoltage = c.PreviousVoltageLevel + gain
		}
	case shared.HalfIncreasingHarvesting:
		gain := shared.StepVoltage * float64(shared.RandInt(1, 100)) / 2.0
		if (c.PreviousVoltageLevel + gain) > shared.MaximumVoltage {
			capacitorVoltage = shared.MaximumVoltage
		} else {
			capacitorVoltage = c.PreviousVoltageLevel + gain
		}
	case shared.QuarterIncreasingHarvesting:
		gain := shared.StepVoltage * float64(shared.RandInt(1, 100)) / 4.0
		if (c.PreviousVoltageLevel + gain) > shared.MaximumVoltage {
			capacitorVoltage = shared.MaximumVoltage
		} else {
			capacitorVoltage = c.PreviousVoltageLevel + gain
		}
	case shared.DecreasingHarvesting:
		gain := shared.StepVoltage * float64(shared.RandInt(1, 100))
		if (c.PreviousVoltageLevel - gain) < 0 {
			capacitorVoltage = 0.0
		} else {
			capacitorVoltage = c.PreviousVoltageLevel - gain
		}
	case shared.HalfDecreasingHarvesting:
		gain := shared.StepVoltage * float64(shared.RandInt(1, 100)) / 2.0
		if (c.PreviousVoltageLevel - gain) < 0 {
			capacitorVoltage = 0.0
		} else {
			capacitorVoltage = c.PreviousVoltageLevel - gain
		}
	case shared.QuarterDecreasingHarvesting:
		gain := shared.StepVoltage * float64(shared.RandInt(1, 100)) / 4.0
		if (c.PreviousVoltageLevel - gain) < 0 {
			capacitorVoltage = 0.0
		} else {
			capacitorVoltage = c.PreviousVoltageLevel - gain
		}
	case shared.ConstantHarvesting:
		capacitorVoltage = shared.StepVoltage * 5 // 2 TODO
	default:
		fmt.Println("Something wrong in the cap behaviour!!", pattern)
		os.Exit(0)
	}
	c.PreviousVoltageLevel = capacitorVoltage

	return capacitorVoltage
}
