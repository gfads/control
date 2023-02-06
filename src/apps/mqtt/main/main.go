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

	var c ops.IController

	// OnOff controllers -- select one
	c = ops.NewController(shared.BasicOnoff, 0.0, 1000)
	c = ops.NewController(shared.DeadZoneOnoff, 0.0, 1000, 10)
	c = ops.NewController(shared.HysteresisOnoff, 0.0, 1000, 100)

	// PID controllers -- select one
	c = ops.NewController(shared.BasicPid, 1.0, 1200, -9600, 0.0, 0.0)  // P Controller
	c = ops.NewController(shared.BasicPid, 1.0, 1200, -9600, 0.01, 0.0) // PI Controller
	c = ops.NewController(shared.BasicPid, 1.0, 1200, -9600, 0.01, 1.0) //
	c = ops.NewController(shared.BasicPid, 1.0, 1200, -9600, 0.01, 1.0)
	c = ops.NewController(shared.SmoothingPid, 1.0, 1200, -9600, 0.01, 0.1)
	c = ops.NewController(shared.IncrementalFormPid, 1.0, 1200, -9600, 0.01, 1.0)
	c = ops.NewController(shared.ErrorSquarePid, 1.0, 1200, -9600, 0.01, 0.1)
	c = ops.NewController(shared.DeadZonePid, 1.0, 1200, -9600, 0.01, 0.1, 6)

	// Gain Scheduling controller (2 set of gains)
	c = ops.NewController(shared.GainScheduling, 0.0, 1000, 1.0, 1.1, 0.0, 6)

	// Gain Scheduling controller
	c = ops.NewController(shared.BasicOnoff, 1.0, 1200)

	// create client and use the controller instance
	client := myclient.NewMyMQTTClient(shared.BrokerAddress, shared.BrokerPort, c)

	// start client
	client.Run()
}
