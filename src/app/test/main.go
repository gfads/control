package main

import (
	"controllers"
	"fmt"
	"shared"
)

func main() {

	// OnOff controllers
	c := controllers.NewController(shared.BASIC_ONOFF, 0.0, 1000)
	c = controllers.NewController(shared.DEAD_ZONE_ONOFF, 0.0, 1000, 10)
	c = controllers.NewController(shared.HYSTERESIS_ONOFF, 0.0, 1000, 100)

	// PID controllers
	c = controllers.NewController(shared.BASIC_PID, 0.0, 1000, 1.0, 1.1, 2.1)
	c = controllers.NewController(shared.SMOOTHING_PID, 0.0, 1000, 1.0, 1.1, 0.1)
	c = controllers.NewController(shared.INCREMENTAL_FORM_PID, 0.0, 1000, 1.0, 1.1, 1.0)
	c = controllers.NewController(shared.ERROR_SQUARE_PID, 0.0, 1000, 1.0, 1.1, 0.0)
	c = controllers.NewController(shared.DEAD_ZONE_PID, 0.0, 1000, 1.0, 1.1, 0.0, 6)

	// Gain Scheduling controller (2 set of gains)
	c = controllers.NewController(shared.GAIN_SCHEDULING, 0.0, 1000, 1.0, 1.1, 0.0, 6)

	fmt.Println(c.Update(11, 10))
}
