package racingcar_test

import (
	"strings"
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

func TestNewGame(t *testing.T) {
	testCases := []struct {
		names []string
	}{
		{names: []string{"test1", "test2"}},
		{names: []string{"test1", "test2", "test3"}},
		{names: []string{"test1", " test2"}},
		{names: []string{"test1", "  test2  "}},
	}
	for _, tc := range testCases {
		game, _ := racingcar.NewGame(tc.names)
		for i, car := range game.Cars {
			assert.Equal(t, strings.TrimSpace(tc.names[i]), car.Name(), "car name: %s", car.Name())
		}
	}
}

func TestNewGameError(t *testing.T) {
	testCases := []struct {
		names []string
	}{
		{names: make([]string, 0)},
		{names: []string{"", "test2"}},
		{names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}},
	}
	for _, tc := range testCases {
		_, err := racingcar.NewGame(tc.names)
		assert.Error(t, err, "error expected %v", tc.names)
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
