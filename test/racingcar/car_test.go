package racingcar_test

import (
	"testing"

	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
)

func TestNewCarNameValidation(t *testing.T) {
	testCases := []struct {
		name        string
		expectedMsg string
	}{
		{name: "", expectedMsg: "empty name"},
		{name: "123456", expectedMsg: "too long name"},
		{name: "abcdef", expectedMsg: "too long name"},
		{name: "한글다섯글자", expectedMsg: "too long name"},
	}
	for _, tc := range testCases {
		_, err := racingcar.NewCar(tc.name)
		assert.EqualError(t, err, tc.expectedMsg, "name: %s, expected : %v, actual : %v", tc.name, tc.expectedMsg, err)
	}
}

func TestCarMoveForward(t *testing.T) {
	car, _ := racingcar.NewCar("test")
	car.MoveForward()
	actualPos := car.Position()
	expectedPos := 1

	assert.Equal(t, expectedPos, actualPos, "moved position: %d", actualPos)
}
