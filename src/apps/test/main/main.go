package main

import (
	"controllers/def/ops"
	"fmt"
	"reflect"
	"shared"
)

func main() {

	// OnOff controllers
	c := ops.NewController(shared.BasicOnoff, 0.0, 1000)
	c = ops.NewController(shared.DeadZoneOnoff, 0.0, 1000, 10)
	c = ops.NewController(shared.HysteresisOnoff, 0.0, 1000, 100)

	// PID controllers
	c = ops.NewController(shared.BasicPid, 0.0, 1000, 1.0, 1.1, 2.1)
	c = ops.NewController(shared.SmoothingPid, 0.0, 1000, 1.0, 1.1, 0.1)
	c = ops.NewController(shared.IncrementalFormPid, 0.0, 1000, 1.0, 1.1, 1.0)
	c = ops.NewController(shared.ErrorSquarePid, 0.0, 1000, 1.0, 1.1, 0.0)
	c = ops.NewController(shared.DeadZonePid, 0.0, 1000, 1.0, 1.1, 0.0, 6)

	// Gain Scheduling controller (2 set of gains)
	c = ops.NewController(shared.GainScheduling, 0.0, 1000, 1.0, 1.1, 0.0, 6)

	fmt.Println(c.Update(11, 10), reflect.TypeOf(c).Elem())
}
