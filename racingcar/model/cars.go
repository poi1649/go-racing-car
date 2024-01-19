package model

type Cars struct {
	Cars []*Car
}

func NewCars(cars []*Car) *Cars {
	return &Cars{Cars: cars}
}

func GenerateCars(names []string) (*Cars, error) {
	cars := make([]*Car, len(names))
	for i, name := range names {
		car, err := NewCar(name)
		if err != nil {
			return nil, err
		}
		cars[i] = car
	}
	return NewCars(cars), nil
}

func (c *Cars) GetWinners() []*Car {
	winners := make([]*Car, 0)
	maxPosition := 0
	for _, car := range c.Cars {
		if car.Position > maxPosition {
			winners = []*Car{car}
			maxPosition = car.Position
			continue
		}
		if car.Position == maxPosition {
			winners = append(winners, car)
		}
	}
	return winners
}
