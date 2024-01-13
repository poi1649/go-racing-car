package test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/thedevluffy/go-racing-car/racingcar"
	"testing"
)

func TestMovableCarMoveForward(t *testing.T) {
	car := Car{Name: "개발자동차", Position: 0}

	movedCar := car.Move()

	assert.Equal(
		t,
		car.Position+1,
		movedCar.Position,
		"Expected car position to be %d, but got %d", car.Position+1, movedCar.Position,
	)
}

func TestMovableCarStop(t *testing.T) {
	car := Car{Name: "개발자동차", Position: 0}

	stoppedCar := car.Stop()

	assert.Equal(
		t,
		car.Position,
		stoppedCar.Position,
		"Expected car position to be %d, but got %d", car.Position+1, stoppedCar.Position,
	)
}
