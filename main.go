package main

import (
	"StrategyPattern/model"
	"fmt"
)

func main() {
	normalVehicle := model.NormalVehicle{}
	normalVehicle.VDrive(normalVehicle.NormalDrive)

	fmt.Println()

	sportVehicle := model.SportVehicle{}
	sportVehicle.VDrive(sportVehicle.SpecialDrive)
}
