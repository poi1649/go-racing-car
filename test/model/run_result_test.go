package model

import (
	"github.com/poi1649/go-racing-car/racingcar/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRunResult(t *testing.T) {
	cars, _ := model.GenerateCars([]string{"test1", "test2"})
	result := model.GenerateRunResult(cars)
	assert.Equal(t, result.CarNameToPosition["test1"], 0)
	assert.Equal(t, result.CarNameToPosition["test2"], 0)
}
