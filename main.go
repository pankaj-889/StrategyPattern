package main

import (
	"StrategyPattern/model"
)

func main() {
	normalVehicle := model.NormalVehicle{}
	vh := model.Vehicle{
		Drive: normalVehicle.NormalDrive,
	}
	vh.VDrive()
	sportVehicle := model.SportVehicle{}
	vh2 := model.Vehicle{
		Drive: sportVehicle.SpecialDrive,
	}
	vh2.VDrive()
}
