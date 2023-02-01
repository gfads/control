package shared

import (
	"math/rand"
)

// MQTT CLIENT
const NODE1_ID = "eui-70b3d57ed005867f"
const NODE2_ID = "eui-70b3d57ed0058955"
const BROKER_ADDRESS = "eu1.cloud.thethings.network"
const BROKER_PORT = 1883
const USER_NAME_MQTT = "floodsensor-appid@ttn"
const USER_PWD_MQTT = "NNSXS.GTSVUEBSLQEQEGFB6M7HACHU7C4IAWSPZCWLZUI.UFN7TKM7UJVP6X63Z6HLU2VCC2T2FATWZUW5ZJMTO4GTMD3A4PFA"

const SAMPLING_CYCLE_SIZE = 50

const SV = 2.7           // Shutoff voltage (page 17) = 2.7 V
const OV = 3.7           // Optimum voltage (page 17) = 3.7 V
const HYSTERISIS = 0.001 // 10 mV

// Controller types
const PID = "PID"
const PI = "PI"
const ONOFF = "OnOff"
const DEAD_ZONE_ONOFF = "OnOffwithDeadZone"
const HYSTERESIS_ONOFF = "OnOffwithHysteresis"
const GAIN_SCHEDULING = "GainScheduling"

// Capacitor behaviour
const INCREASING = 0
const HALF_INCREASING = 1
const QUARTER_INCREASING = 2
const DECREASING = 3
const HALF_DECREASING = 4
const QUARTER_DECREASING = 5
const CONSTANT = 6
const RANDOM = 7

const MAXIMUM_VOLTAGE = 100.0
const STEP_VOLTAGE = 0.001 // in Volts

const ADAPTATION_CYCLES = 131

const MINIMUM_TASK_EXECUTION_RATE = 1    // page 17
const MAXIMUM_TASK_EXECUTION_RATE = 1200 // page 17

const INITIAL_VOLTAGE = 3.3 // 3.3 V - page 19
const INITIAL_RATE = 1

const AVERAGE_CONSUMPTION = 0.01

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

const BASIC_PID = "Basic"
const SMOOTHING_PID = "SmoothingDerivative"
const INCREMENTAL_FORM_PID = "IncrementalForm"
const ERROR_SQUARE_PID = "ErrorSquare"
const NONE_PID = "None"
const DEAD_ZONE_PID = "DeadZonePID"
