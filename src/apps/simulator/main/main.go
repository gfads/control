package main

import (
	"astar/algorithm"
	"astar/cap"
	"controllers/def/ops"
	"fmt"
	"math"
	"shared"
)

func main() {

	// AsTAR elements
	rold := 1                     // previous rate
	vnew := shared.InitialVoltage // initial capacitor voltage
	astar := algorithm.AsTAR{}

	// OnOff controllers
	controller := ops.NewController(shared.BasicOnoff, 0.0, 1000)
	controller = ops.NewController(shared.DeadZoneOnoff, 0.0, 1000, 10)
	controller = ops.NewController(shared.HysteresisOnoff, 0.0, 1000, 100)

	// PID controllers
	controller = ops.NewController(shared.BasicPid, 0.0, 1000, 1.0, 1.1, 2.1)
	controller = ops.NewController(shared.SmoothingPid, 0.0, 1000, 1.0, 1.1, 0.1)
	controller = ops.NewController(shared.IncrementalFormPid, 0.0, 1000, 1.0, 1.1, 1.0)
	controller = ops.NewController(shared.ErrorSquarePid, 0.0, 1000, 1.0, 1.1, 0.0)
	controller = ops.NewController(shared.DeadZonePid, 0.0, 1000, 1.0, 1.1, 0.0, 6)

	// Gain Scheduling controller (2 set of gains)
	controller = ops.NewController(shared.GainScheduling, 0.0, 1000, 1.0, 1.1, 0.0, 6)

	capacitor := cap.Capacitor{Pattern: shared.RandomHarvesting, PreviousVoltageLevel: shared.InitialVoltage}
	for i := 0; i < shared.AdaptationCycles; i++ {

		// obtain the cap voltage level
		vold := capacitor.PreviousVoltageLevel
		vnew = capacitor.GetCapacitorVoltageLevel() // simulator
		//vnew = capacitor.GetCapacitorVoltageLevelFromFile()

		// update task rate -- AsTAR and Controller
		rnewAstar, _ := astar.Update(vnew, vold, rold) // second outpur parameter not used
		rnewController := controller.Update(3.7, vnew)

		fmt.Printf("%.2f;%d;%d\n", vnew, rnewAstar, int(math.Round(rnewController)))

		// update rate
		rold = rnewAstar
	}
	fmt.Scanln()
}
