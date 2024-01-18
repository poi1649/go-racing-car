package test

import (
	"testing"

	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
)

func TestPlayTurn(t *testing.T) {
	testCases := []struct {
		moveStrategy racingcar.MoveStrategy
		expectedPos  int
	}{
		{func() bool { return true }, 1},
		{func() bool { return false }, 0},
	}
	for _, tc := range testCases {
		car, _ := racingcar.NewCar("test")
		racingcar.PlayTurn(car, tc.moveStrategy)
		actualPos := car.Position()

		assert.Equal(t, tc.expectedPos, actualPos, "moved position: %d", actualPos)
	}
}

func TestGamePlayTurn(t *testing.T) {
	testMoveStrategy := getTestStrategy()
	game, _ := racingcar.NewGame([]string{"test1", "test2"})

	game.PlayTurn(testMoveStrategy)

	actualPos1 := game.Cars[0].Position()
	actualPos2 := game.Cars[1].Position()
	expectedPos1 := 1
	expectedPos2 := 0

	assert.Equal(t, expectedPos1, actualPos1, "car1 moved position: %d", actualPos1)
	assert.Equal(t, expectedPos2, actualPos2, "car2 moved position: %d", actualPos2)
}

// test streategy : true, false, true, false, ... 번갈아가며 반환
func getTestStrategy() func() bool {
	isMovable := true
	return func() bool {
		result := isMovable
		isMovable = !isMovable
		return result
	}
}
