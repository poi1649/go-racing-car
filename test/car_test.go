package test

import (
	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCar(t *testing.T) {
	car, err := racingcar.NewCar("test")
	assert.NotNilf(t, car, "car should not be nil")
	assert.Nilf(t, err, "err should be nil")
}
