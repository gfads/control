package ops

import (
	gainscheduling "controllers/gain"
	"controllers/onoff/basic"
	deadzoneonff "controllers/onoff/deadzone"
	hysteresisonoff "controllers/onoff/hysteresis"
	"controllers/pid/basic"
	deadzonepid "controllers/pid/deadzone"
	errorsquarepid "controllers/pid/errorsquare"
	incrementalpid "controllers/pid/incremental"
	smoothingpid "controllers/pid/smoothing"
	"fmt"
	"os"
	"shared"
)

type IController interface {
	Initialise(...float64)
	Update(...float64) float64
}

func NewController(typeName string, p ...float64) IController {

	switch typeName {
	case shared.BASIC_ONOFF:
		c := onoffbasic.Controller{}
		c.Initialise(p...)
		return &c
	case shared.DEAD_ZONE_ONOFF:
		c := deadzoneonff.Controller{}
		c.Initialise(p...)
		return &c
	case shared.HYSTERESIS_ONOFF:
		c := hysteresisonoff.Controller{}
		c.Initialise(p...)
		return &c
	case shared.BASIC_PID:
		c := basicpid.Controller{}
		c.Initialise(p...)
		return &c
	case shared.SMOOTHING_PID:
		c := smoothingpid.Controller{}
		c.Initialise(p...)
		return &c
	case shared.INCREMENTAL_FORM_PID:
		c := incrementalpid.Controller{}
		c.Initialise(p...)
		return &c
	case shared.DEAD_ZONE_PID:
		c := deadzonepid.Controller{}
		c.Initialise(p...)
		return &c
	case shared.ERROR_SQUARE_PID:
		c := errorsquarepid.Controller{}
		c.Initialise(p...)
		return &c
	case shared.GAIN_SCHEDULING:
		c := gainscheduling.Controller{}
		c.Initialise(p...)
		return &c
	default:
		fmt.Println("Error: Controller type ´", typeName, "´ is unknown!")
		os.Exit(0)
	}

	return *new(IController)
}
