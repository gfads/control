package info

type Controller struct {
	TypeName string
	Kp       float64
	Ki       float64
	Kd       float64

	Min float64
	Max float64

	Integrator             float64
	SumPreviousErrors      float64
	PreviousOut            float64
	PreviousError          float64
	PreviousPreviousError  float64
	PreviousDifferentiator float64
	DeadZone               float64
	HysteresisBand         float64
	Out                    float64
}
