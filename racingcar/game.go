package racingcar

import (
	"fmt"
	"strings"
)

const (
	minParticipants     = 1
	maxParticipants     = 20
	allowedMinTurnCount = 1
	allowedMaxTurnCount = 100
)

type RaceCars []Car

type Game struct {
	Cars             RaceCars
	maxTurnCount     int
	currentTurnCount int
}

func NewGame(carNames []string, maxTurnCount int) (*Game, error) {
	cars, err := initCars(carNames)
	if err != nil {
		return nil, err
	}

	if !validateGameCount(maxTurnCount) {
		return nil, fmt.Errorf("시도 횟수는 %d이상 %d이하의 숫자만 가능합니다.", allowedMinTurnCount, allowedMaxTurnCount)
	}

	return &Game{
		Cars:             cars,
		maxTurnCount:     maxTurnCount,
		currentTurnCount: 0,
	}, nil
}

func initCars(carNames []string) ([]Car, error) {
	if !validateNumberOfCar(len(carNames)) {
		return nil, fmt.Errorf("자동차의 수는 %d대 이상 %d대 이하만 가능합니다.", minParticipants, maxParticipants)
	}

	cars := make([]Car, len(carNames))
	for i, name := range carNames {
		car, err := NewCar(strings.TrimSpace(name))
		if err != nil {
			return nil, err
		}
		cars[i] = *car
	}

	return cars, nil
}

func validateNumberOfCar(numberOfCar int) bool {
	return numberOfCar >= minParticipants && numberOfCar <= maxParticipants
}

func validateGameCount(count int) bool {
	return count >= allowedMinTurnCount && count <= allowedMaxTurnCount
}

func (game *Game) IsEnd() bool {
	return game.currentTurnCount == game.maxTurnCount
}

func (game *Game) PlayTurn(moveStrategy MoveStrategy) error {
	if game.IsEnd() {
		return fmt.Errorf("게임이 종료되었습니다.")
	}
	for i := range game.Cars {
		PlayTurn(&game.Cars[i], moveStrategy)
	}
	game.currentTurnCount++
	return nil
}

func PlayTurn(car *Car, isMovable MoveStrategy) {
	if isMovable() {
		car.MoveForward()
	}
}

func (game *Game) Winners() []Car {
	var winners []Car
	maxPosition := 0
	for _, car := range game.Cars {
		if car.Position() > maxPosition {
			maxPosition = car.Position()
			winners = []Car{car}
		} else if car.Position() == maxPosition {
			winners = append(winners, car)
		}
	}
	return winners
}
