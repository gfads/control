package algorithm

import (
	"shared"
)

type AsTAR struct {
}

func (a AsTAR) Update(vnew float64, vold float64, rold int) (int, float64) {
	rnew := 0
	getnew := 0.0

	if vnew < shared.SV { // The system is in Shut-off Voltage state, task is stopped
		rnew = 0.0
	} else if vnew < (shared.OV - shared.HYSTERESIS) { // The system is in Low-voltage state, apply AIMD
		if vnew > vold {
			rnew = rold + 1
		} else {
			rnew = rold / 2
		}
	} else if vnew > (shared.OV + shared.HYSTERESIS) { // The system is in High Voltage state, apply MIAD
		if vnew < vold {
			rnew = rold - 1
		} else {
			rnew = rold * 2
		}
	} else { // The system is at Optimum Voltage state, take no action
		rnew = rold
	}

	// final check of rnew
	if rnew < shared.MinimumTaskExecutionRate {
		rnew = shared.MinimumTaskExecutionRate
	}
	if rnew > shared.MaximumTaskExecutionRate {
		rnew = shared.MaximumTaskExecutionRate
	}

	if rnew != 0 {
		getnew = 1.0 / float64(rnew)
	} else {
		getnew = 0
	}

	return rnew, getnew
}
