package racingcar_test

import (
	"testing"

	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
)

func TestNewCar(t *testing.T) {
	testValid := func(name string) func(t *testing.T) {
		return func(t *testing.T) {
			car, err := racingcar.NewCar(name)
			if assert.NoError(t, err) {
				assert.Equal(t, name, car.Name())
				assert.Equal(t, 0, car.Position())
			}
		}
	}
	testError := func(name string, expectedMsg string) func(t *testing.T) {
		return func(t *testing.T) {
			_, err := racingcar.NewCar(name)
			assert.EqualError(t, err, expectedMsg, "name: %s, expected : %v, actual : %v", name, expectedMsg, err)
		}
	}

	testCases := map[string]func(t *testing.T){
		"valid name 'a'":       testValid("a"),
		"valid name 'test'":    testValid("test"),
		"valid name '자동차'":     testValid("자동차"),
		"valid name '12345'":   testValid("12345"),
		"invlid name ''":       testError("", "empty name"),
		"invlid name 'abcdef'": testError("abcdef", "too long name"),
		"invlid name '한글다섯글자'": testError("한글다섯글자", "too long name"),
	}
	for name, testFunc := range testCases {
		t.Run(name, testFunc)
	}
}

func TestCarMoveForward(t *testing.T) {
	car, _ := racingcar.NewCar("test")
	car.MoveForward()
	actualPos := car.Position()
	expectedPos := 1

	assert.Equal(t, expectedPos, actualPos, "moved position: %d", actualPos)
}
