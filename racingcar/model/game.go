package model

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

func (g *Game) Run() (*RunResult, error) {
	for _, car := range g.Cars.Cars {
		fuel, _ := g.FuelGenerator.Fuel()
		err := car.Move(fuel)
		if err != nil {
			return nil, err
		}
	}
	return GenerateRunResult(g.Cars), nil
}

func (g *Game) GetWinnersName() []string {
	winners := g.Cars.GetWinners()
	winnersName := make([]string, len(winners))
	for i, winner := range winners {
		winnersName[i] = winner.Name
	}
	return winnersName
}
