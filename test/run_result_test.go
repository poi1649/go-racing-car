package test

import (
	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRunResult(t *testing.T) {
	cars, _ := racingcar.GenerateCars([]string{"test1", "test2"})
	result := racingcar.GenerateRunResult(cars)
	assert.Equal(t, result.CarNameToPosition["test1"], 0)
	assert.Equal(t, result.CarNameToPosition["test2"], 0)
}
