package main

import (
	"app/mqtt/myclient"
	"controllers"
	"fmt"
	"shared"
)

func main() {

	//controller := controller.NewController(shared.PID, shared.DEAD_ZONE_PID, -9600, 0.01, 1.0, 2.7/10)
	//controller := controller.NewController(shared.PID, shared.INCREMENTAL_FORM_PID, -9600, 0.01, 1.0)
	//controller := controller.NewController(shared.PID, shared.ERROR_SQUARE_PID, -9600, 0.01, 1.0)
	//controller := controller.NewController(shared.PID, shared.SMOOTHING_PID, -9600, 0.01, 1.0)
	//controller := controller.NewController(shared.PID, shared.BASIC_PID, -9600, 0.0, 0.0) // P Controller
	//controller := controller.NewController(shared.PID, shared.BASIC_PID, -9600, 0.01, 0.0) // PI Controller
	//controller := controller.NewController(shared.PID, shared.BASIC_PID, -9600, 0.01, 1.0) // PID Controller
	//controller := controller.NewController(shared.ONOFF, "", 1.0, 1200.0) // OnOff controller
	//controller := controller.NewController(shared.HYSTERESIS_ONOFF, "", 1.0, 1200.0, 2.7/16.0) // 25% of 2.7 Hysteresis zone controller
	//controller := controller.NewController(shared.DEAD_ZONE_ONOFF, "", 1.0, 1200.0, 2.7/16.0) // Dead zone controller
	//controller := controller.NewController(shared.GAIN_SCHEDULING, shared.BASIC_PID, 0.0, 0.0, 0.0) // Gain Scheduling PID adaptive control

	myController := controllers.NewController(shared.ONOFF, "", 1.0, 1200.0)
	myMqqtClient := myclient.CreateMyMQTTClient(shared.BROKER_ADDRESS, shared.BROKER_PORT, myController)

	myMqqtClient.Run()

	//capacitor := cap.Capacitor{Pattern: shared.RANDOM, PreviousVoltageLevel: shared.INITIAL_VOLTAGE}
	//for i := 0; i < shared.ADAPTATION_CYCLES; i++ {
	// obtain the cap voltage level
	//vold := capacitor.PreviousVoltageLevel
	//vnew = capacitor.GetCapacitorVoltageLevel() // simulator
	//vnew = capacitor.GetCapacitorVoltageLevelFromFile()
	//vnew = capacitor.GetCapacitorVoltagelevelFromTTN()
	//taskrate := controller(vnew)
	//ttnclient.publish(taskrate)
	//vnew = capacitor.GetCapacitorVoltagelevelFromInfluxDB()
	//vnew =

	// update rate
	//rnewAstar, _ := astar.Update(vnew, vold, rold)
	//rnewController := controller.Update(3.7, vnew)

	//fmt.Printf("%.2f;%d;%d\n", vnew, rnewAstar, int(math.Round(rnewController)))

	// update rate
	//rold = rnewAstar
	//}

	fmt.Scanln()
}
