package racingcar_test

import (
	"testing"

	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	testValid := func(names []string, trycount int) func(t *testing.T) {
		return func(t *testing.T) {
			game, err := racingcar.NewGame(names, trycount)
			if assert.NoError(t, err) {
				assert.Equal(t, len(names), len(game.Cars), "generated car count: %d", len(game.Cars))
			}
		}
	}
	testError := func(names []string, trycount int, expectedMsg string) func(t *testing.T) {
		return func(t *testing.T) {
			_, err := racingcar.NewGame(names, trycount)
			assert.EqualError(t, err, expectedMsg, "names: %v, trycount: %d, expected : %v, actual : %v", names, trycount, expectedMsg, err)
		}
	}

	testCases := map[string]func(t *testing.T){
		"valid names ['a'], trycount 1":        testValid([]string{"a", "b"}, 1),
		"valid names ['a', 'b'], trycount 1":   testValid([]string{"a", "b"}, 1),
		"valid names ['a', 'b'], trycount 100": testValid([]string{"a", "b"}, 100),

		"invalid names [], trycount 1":                     testError(make([]string, 0), 1, "자동차의 수는 1대 이상 20대 이하만 가능합니다."),
		"invalid names ['1', '2', ...,  '21'], trycount 1": testError([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}, 1, "자동차의 수는 1대 이상 20대 이하만 가능합니다."),
		"invvalid names ['a', 'b'], trycount -1":           testError([]string{"a", "b"}, -1, "시도 횟수는 1이상 100이하의 숫자만 가능합니다."),
		"invvalid names ['a', 'b'], trycount 0":            testError([]string{"a", "b"}, 0, "시도 횟수는 1이상 100이하의 숫자만 가능합니다."),
		"invvalid names ['a', 'b'], trycount 101":          testError([]string{"a", "b"}, 101, "시도 횟수는 1이상 100이하의 숫자만 가능합니다."),
	}
	for name, tc := range testCases {
		t.Run(name, tc)
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
