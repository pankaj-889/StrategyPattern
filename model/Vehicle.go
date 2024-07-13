package model

import "StrategyPattern/Strategy"

type Vehicle struct {
	Drive Strategy.IDriveStrategy
}

func (v *Vehicle) VDrive() {
	v.Drive.Drive()
}
