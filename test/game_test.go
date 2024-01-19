package test

import (
	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	game, _ := racingcar.NewGame([]string{"test1", "test2"}, racingcar.NewFixedFuelGenerator([]int{1, 2, 3}))
	assert.Equal(t, game.Cars.Cars[0].Name, "test1")
}
