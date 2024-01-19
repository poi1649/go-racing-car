package racingcar

type Game struct {
	Cars          *Cars
	FuelGenerator FuelGenerator
}

func NewGame(carNames []string, fuelGenerator FuelGenerator) (*Game, error) {
	cars, err := GenerateCars(carNames)
	if err != nil {
		return nil, err
	}
	return &Game{Cars: cars, FuelGenerator: fuelGenerator}, nil
}
