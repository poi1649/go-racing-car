package racingcar

type Cars struct {
	Cars []*Car
}

func NewCars(cars []*Car) *Cars {
	return &Cars{Cars: cars}
}
