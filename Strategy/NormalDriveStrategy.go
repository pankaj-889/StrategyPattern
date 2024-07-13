package Strategy

type NormalDriveStrategy struct {
}

func (n *NormalDriveStrategy) NewNormalDriveStrategy() *NormalDriveStrategy {
	return &NormalDriveStrategy{}
}

func (n *NormalDriveStrategy) Drive() {
	println("Normal Drive")
}
