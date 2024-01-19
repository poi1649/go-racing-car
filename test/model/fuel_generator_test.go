package model

import (
	"github.com/poi1649/go-racing-car/racingcar/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFuelGenerator(t *testing.T) {
	fuels := []int{1, 2, 3}
	generator := model.NewFixedFuelGenerator(fuels)
	fuel, _ := generator.Fuel()
	assert.Equal(t, 1, fuel)
}

func TestNewFuelGeneratorWithOverLengthFuel(t *testing.T) {
	fuels := []int{1, 2, 3}
	generator := model.NewFixedFuelGenerator(fuels)
	generator.Fuel()
	generator.Fuel()
	generator.Fuel()
	_, err := generator.Fuel()
	assert.Equal(t, "제공된 연료 길이보다 많은 연료를 요청했습니다", err.Error())
}
