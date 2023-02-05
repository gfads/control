package harvesting

import (
	"fmt"
	"os"
	"shared"
)

func HarvestedEnergy(previousVoltage float64, pattern int) float64 {
	harvestedVoltage := 0.0

	if pattern == shared.RandomHarvesting {
		pattern = shared.RandInt(0, shared.RandomHarvesting)
	}

	switch pattern {
	case shared.NoHarvesting:
		harvestedVoltage = 0.0
	case shared.IncreasingHarvesting:
		if (previousVoltage + shared.StepVoltage) > shared.MaximumVoltage {
			harvestedVoltage = shared.MaximumVoltage
		} else {
			harvestedVoltage = previousVoltage + shared.StepVoltage
		}
	case shared.DecreasingHarvesting:
		if (previousVoltage - shared.StepVoltage) < 0 {
			harvestedVoltage = 0.0
		} else {
			harvestedVoltage = previousVoltage - shared.StepVoltage
		}
	case shared.ConstantHarvesting:
		harvestedVoltage = shared.StepVoltage * 5 // Variable
	default:
		fmt.Println("Something wrong in the harvest system!!")
		os.Exit(0)
	}
	return harvestedVoltage
}
