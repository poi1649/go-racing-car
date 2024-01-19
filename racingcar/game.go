package racingcar

import (
	"fmt"
	"math/rand"
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

const (
	defaultMoveThreshod = 4
	defaultMoveNumRange = 10
)

type MoveStrategy func() bool

func DefalutMoveStrategy() bool {
	return rand.Intn(defaultMoveNumRange) >= defaultMoveThreshod
}

func PlayTurn(car *Car, isMovable MoveStrategy) {
	if isMovable() {
		car.MoveForward()
	}
}

// 위의 함수를 바는 방법과 아래의 인터페이스와 빈 struct를 사용하는 방법중 고민
// 위의 함수를 이용한 방법은 좀 더 간단하고 단순하게 사용할 수 있다는 장점이 있는것고
// 아래의 인터페이스를 이용한 방법은 좀 더 명시적이고, 확장성이 있는것 같습니다.
// 이동에 대한 더 복잡한 전략이 생길 여지가 있다면 아래의 방법이 좋을 것 같지만 일단 간단하게 미션진행을 위해 위의 방법으로 구현해보았습니다.
// 의견이 궁금해서 주석처리하고 남겨두었습니다. 의견 있으면 공유하고싶어요!!
/*
type MoveStrategy interface {
	IsMovable() bool
}

type DefaultStrategy struct{}

func (DefaultStrategy) IsMovable() bool {
	return rand.Intn(10) >= defaultMoveThresho
}

func playTurn(car *Car, moveStrategy MoveStrategy) {
	if moveStrategy.IsMovable() {
		car.MoveForward()
	}
}
*/
