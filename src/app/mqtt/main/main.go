package main

import (
	"app/mqtt/myclient"
	"controllers"
	"shared"
)

func main() {

	// On Off controllers
	//c := controllers.NewController(shared.BASIC_ONOFF,1.0, 1200.0)
	//c := controllers.NewController(shared.HYSTERESIS_ONOFF,  1.0, 1200.0, 2.7/16.0)
	//c := controllers.NewController(shared.DEAD_ZONE_ONOFF,  1.0, 1200.0, 2.7/16.0)

	// PID controllers
	//c := controllers.NewController(shared.BASIC_PID, -9600, 0.0, 0.0) // P Controller
	//c := controllers.NewController(shared.BASIC_PID, -9600, 0.01, 0.0) // PI Controller
	//c := controllers.NewController(shared.BASIC_PID, -9600, 0.01, 1.0) // PID Controllerc = controllers.NewController(shared.DEAD_ZONE_PID, -9600, 0.01, 1.0, 2.7/10)
	//c := controllers.NewController(shared.INCREMENTAL_FORM_PID, -9600, 0.01, 1.0)
	//c := controllers.NewController(shared.ERROR_SQUARE_PID, -9600, 0.01, 1.0)
	//c := controllers.NewController(shared.SMOOTHING_PID, -9600, 0.01, 1.0)

	// Gain Scheduling controller
	c := controllers.NewController(shared.GAIN_SCHEDULING, 0.0, 0.0, 0.0)

	// create client and use the controller instance
	client := myclient.NewMyMQTTClient(shared.BROKER_ADDRESS, shared.BROKER_PORT, c)

	// start client
	client.Run()
}
