package main

import (
	"app/mqtt/myclient"
	"controllers"
	"shared"
)

func main() {

	// instantiate a controller

	//c := controller.NewController(shared.PID, shared.DEAD_ZONE_PID, -9600, 0.01, 1.0, 2.7/10)
	//c := controller.NewController(shared.PID, shared.INCREMENTAL_FORM_PID, -9600, 0.01, 1.0)
	//c := controller.NewController(shared.PID, shared.ERROR_SQUARE_PID, -9600, 0.01, 1.0)
	//c := controller.NewController(shared.PID, shared.SMOOTHING_PID, -9600, 0.01, 1.0)
	//c := controller.NewController(shared.PID, shared.BASIC_PID, -9600, 0.0, 0.0) // P Controller
	//c := controller.NewController(shared.PID, shared.BASIC_PID, -9600, 0.01, 0.0) // PI Controller
	//c := controller.NewController(shared.PID, shared.BASIC_PID, -9600, 0.01, 1.0) // PID Controller
	//c := controller.NewController(shared.ONOFF, "", 1.0, 1200.0) // OnOff controller
	//c := controller.NewController(shared.HYSTERESIS_ONOFF, "", 1.0, 1200.0, 2.7/16.0) // 25% of 2.7 Hysteresis zone controller
	//c := controller.NewController(shared.DEAD_ZONE_ONOFF, "", 1.0, 1200.0, 2.7/16.0) // Dead zone controller
	//c := controller.NewController(shared.GAIN_SCHEDULING, shared.BASIC_PID, 0.0, 0.0, 0.0) // Gain Scheduling PID adaptive control

	c := controllers.NewController(shared.BASIC_ONOFF, 1.0, 1200.0)

	// create client and use the controller instance
	client := myclient.NewMyMQTTClient(shared.BROKER_ADDRESS, shared.BROKER_PORT, c)

	// start client
	client.Run()
}
