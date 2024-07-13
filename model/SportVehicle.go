package model

import "StrategyPattern/Strategy"

type SportVehicle struct {
	SpecialDrive *Strategy.SpecialDriveStrategy
}

func (s *SportVehicle) NewSportVehicle() *SportVehicle {
	return &SportVehicle{
		SpecialDrive: &Strategy.SpecialDriveStrategy{},
	}
}
