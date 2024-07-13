package Strategy

type SpecialDriveStrategy struct {
}

func (s *SpecialDriveStrategy) NewSpecialDriveStrategy() *SpecialDriveStrategy {
	return &SpecialDriveStrategy{}
}

func (s *SpecialDriveStrategy) Drive() {
	println("Special Drive")
}
