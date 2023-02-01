package controllers

import (
	gain "controllers/gainschedulling"
	"controllers/onoff"
	"controllers/onoffdeadzone"
	"controllers/onoffhysteresis"
	"controllers/pid"
	"fmt"
	"os"
	"shared"
)

type IController interface {
	Update(...float64) float64
	Initialise(string, ...float64)
	Reconfigure(...float64)
	SetKP(float64)
	GetKP() float64
	SetGains(float64, float64, float64)
}

func NewController(typeName string, pidType string, p ...float64) IController {
	switch typeName {
	case shared.ONOFF:
		min := p[0]
		max := p[1]
		c := onoff.OnOffController{}
		c.Initialise("", min, max)
		return &c
	case shared.DEAD_ZONE_ONOFF:
		min := p[0]
		max := p[1]
		band := p[2]
		c := onoffdeadzone.OnOffDeadZoneController{}
		c.Initialise("", min, max, band)
		return &c
	case shared.HYSTERESIS_ONOFF:
		min := p[0]
		max := p[1]
		band := p[2]
		c := onoffhysteresis.OnOffwithHysteresisController{}
		c.Initialise("", min, max, band)
		return &c
	case shared.PID:
		kp := p[0]
		ki := p[1]
		kd := p[2]
		limMin := p[3]
		limMax := p[4]
		deadZone := 0.0
		if pidType == shared.DEAD_ZONE_PID {
			deadZone = p[3]
		}
		c := pid.PIDController{}
		c.Initialise(pidType, kp, ki, kd, float64(limMin), float64(limMax), deadZone)
		return &c
	case shared.GAIN_SCHEDULING:
		kp := p[0]
		ki := p[1]
		kd := p[2]
		limMin := p[3]
		limMax := p[4]
		c := gain.GainPIDController{}
		c.Initialise(pidType, kp, ki, kd, float64(limMin), float64(limMax)) // TODO pidType
		return &c
	case shared.PI:
		kp := p[0]
		ki := p[1]
		kd := p[2]
		limMin := p[3]
		limMax := p[4]
		c := pid.PIDController{}
		c.Initialise(pidType, kp, ki, kd, float64(limMin), float64(limMax))
		return &c
	default:
		fmt.Println("Error: Controller type ´", typeName, "´ is invalid!")
		os.Exit(0)
	}

	return *new(IController)
}

/*
func Update(c IController, r float64, y float64) float64 {
	return c.Update(r, y)
}

func Initialise(c IController, controllerType string, p ...float64) {
	c.Initialise(controllerType, p)
}

func Reconfigure(c IController) {
	c.Reconfigure()
}

func SetKP(c IController, kp float64) {
	c.SetKP(kp)
}

func GetKP(c IController) float64 {
	return c.GetKP()
}

func SetGains(c IController, kp, ki, kd float64) {
	c.SetGains(kp, ki, kd)
}
*/
