package model

import (
	"github.com/poi1649/go-racing-car/racingcar/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	game, _ := model.NewGame([]string{"test1", "test2"}, model.NewFixedFuelGenerator([]int{1, 2, 3}))
	assert.Equal(t, game.Cars.Cars[0].Name, "test1")
}

func TestGameRun(t *testing.T) {
	game, _ := model.NewGame([]string{"test1", "test2"}, model.NewFixedFuelGenerator([]int{1, 4}))
	result, _ := game.Run()
	assert.Equal(t, result.CarNameToPosition["test1"], 0)
}

func TestGameGetWinnersName(t *testing.T) {
	game, _ := model.NewGame([]string{"test1", "test2"}, model.NewFixedFuelGenerator([]int{1, 4}))
	game.Run()
	assert.Equal(t, []string{"test2"}, game.GetWinnersName())
}

func TestGameGetWinnersNameWithSamePosition(t *testing.T) {
	game, _ := model.NewGame([]string{"test1", "test2"}, model.NewFixedFuelGenerator([]int{4, 4}))
	game.Run()
	assert.Equal(t, []string{"test1", "test2"}, game.GetWinnersName())
}
