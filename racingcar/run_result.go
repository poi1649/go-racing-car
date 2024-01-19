package racingcar

type RunResult struct {
	CarNameToPosition map[string]int
}

func NewRunResult(carNameToPosition map[string]int) *RunResult {
	return &RunResult{CarNameToPosition: carNameToPosition}
}

func GenerateRunResult(cars *Cars) *RunResult {
	carNameToPosition := make(map[string]int)
	for _, car := range cars.Cars {
		carNameToPosition[car.Name] = car.Position
	}
	return NewRunResult(carNameToPosition)
}
