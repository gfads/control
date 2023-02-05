package dev

import "shared"

func ConsumedEnergy(rold int, GETnew float64) float64 {

	//return float64(rold) * shared.AVERAGE_CONSUMPTION * GETnew
	return float64(rold) * shared.AverageConsumption
}
