package racingcar_test

import (
	"testing"

	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
)

func TestNewCarNameValidation(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{""},
		{"12345"},
		{"한글다섯글자"},
	}
	for _, tc := range testCases {
		_, err := racingcar.NewCar(tc.name)
		assert.Error(t, err, "name: %s", tc.name)
	}
}

func TestCarMoveForward(t *testing.T) {
	car, _ := racingcar.NewCar("test")
	car.MoveForward()
	actualPos := car.Position()
	expectedPos := 1

	assert.Equal(t, expectedPos, actualPos, "moved position: %d", actualPos)
}
