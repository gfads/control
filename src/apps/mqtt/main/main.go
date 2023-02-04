/*********************************************************************************
Author: Nelson S Rosa
Description: This program shows an example of using an MQTT client with a controller. The client
             can use several types of controllers, e.g., On-Off, PID, and Gain
             scheduling.
Date: 04/02/2023
*********************************************************************************/

package main

import (
	"apps/mqtt/myclient"
	"controllers/def/ops"
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
	c := ops.NewController(shared.BasicOnoff, 1.0, 1200)

	// create client and use the controller instance
	client := myclient.NewMyMQTTClient(shared.BrokerAddress, shared.BrokerPort, c)

	// start client
	client.Run()
}
