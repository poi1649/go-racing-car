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
		names    []string
		trycount int
	}{
		{names: []string{"test1", "test2"}, trycount: 1},
		{names: []string{"test1", "test2", "test3"}, trycount: 1},
		{names: []string{"test1", " test2"}, trycount: 1},
		{names: []string{"test1", "  test2  "}, trycount: 1},
	}
	for _, tc := range testCases {
		game, _ := racingcar.NewGame(tc.names, tc.trycount)
		for i, car := range game.Cars {
			assert.Equal(t, strings.TrimSpace(tc.names[i]), car.Name(), "car name: %s", car.Name())
		}
	}
}

func TestNewGameError(t *testing.T) {
	testCases := []struct {
		names    []string
		trycount int
	}{
		{names: make([]string, 0), trycount: 1},
		{names: []string{"", "test2"}, trycount: 1},
		{names: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}, trycount: 1},
		{names: []string{"test1", "test2"}, trycount: -1},
		{names: []string{"test1", "test2"}, trycount: 0},
		{names: []string{"test1", "test2"}, trycount: 101},
	}
	for _, tc := range testCases {
		_, err := racingcar.NewGame(tc.names, tc.trycount)
		assert.Error(t, err, "error expected %v, %d", tc.names, tc.trycount)
	}
}

func TestIsEnd(t *testing.T) {
	testCases := []struct {
		names    []string
		trycount int
	}{
		{names: []string{"test1", "test2"}, trycount: 1},
		{names: []string{"test1", "test2"}, trycount: 2},
		{names: []string{"test1", "test2"}, trycount: 3},
	}
	for _, tc := range testCases {
		game, _ := racingcar.NewGame(tc.names, tc.trycount)
		for i := 0; i < tc.trycount; i++ {
			assert.False(t, game.IsEnd(), "game should not end")
			game.PlayTurn(racingcar.DefalutMoveStrategy)
		}
		assert.True(t, game.IsEnd(), "game should end")
	}
}

func TestGamePlayTurn(t *testing.T) {
	testMoveStrategy := getTestAlternateStrategy()
	game, _ := racingcar.NewGame([]string{"test1", "test2"}, 1)

	game.PlayTurn(testMoveStrategy)

	actualPos1 := game.Cars[0].Position()
	actualPos2 := game.Cars[1].Position()
	expectedPos1 := 1
	expectedPos2 := 0

	assert.Equal(t, expectedPos1, actualPos1, "car1 moved position: %d", actualPos1)
	assert.Equal(t, expectedPos2, actualPos2, "car2 moved position: %d", actualPos2)
}

// test strategy : true, false, true, false, ... 번갈아가며 반환
func getTestAlternateStrategy() func() bool {
	isMovable := true
	return func() bool {
		result := isMovable
		isMovable = !isMovable
		return result
	}
}

func TestWinners(t *testing.T) {
	testCases := []struct {
		names         []string
		strategyFucn  func() bool
		expectedNames []string
	}{
		{
			names:         []string{"test1", "test2"},
			strategyFucn:  getTestAlternateStrategy(),
			expectedNames: []string{"test1"},
		},
		{
			names:         []string{"test1", "test2"},
			strategyFucn:  getTestTrueStrategy(),
			expectedNames: []string{"test1", "test2"},
		},
	}
	for _, tc := range testCases {
		game, _ := racingcar.NewGame(tc.names, 2)
		for i := 0; i < 2; i++ {
			game.PlayTurn(tc.strategyFucn)
		}
		winners := game.Winners()
		winnerNames := make([]string, len(winners))
		for i, winner := range winners {
			winnerNames[i] = winner.Name()
		}

		assert.ElementsMatch(t, tc.expectedNames, winnerNames, "winners: %v", game.Winners())
	}
}

// test strategy to return true only
func getTestTrueStrategy() func() bool {
	return func() bool {
		return true
	}
}
