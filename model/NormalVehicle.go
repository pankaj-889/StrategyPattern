package model

import "StrategyPattern/Strategy"

type NormalVehicle struct {
	NormalDrive *Strategy.NormalDriveStrategy
}

func (n *NormalVehicle) NewNormalVehicle() *NormalVehicle {
	return &NormalVehicle{
		NormalDrive: &Strategy.NormalDriveStrategy{},
	}
}
