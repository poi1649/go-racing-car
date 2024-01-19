package test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/thedevluffy/go-racing-car/racingcar"
	"testing"
)

func TestMovableCarMoveForward(t *testing.T) {
	car := Car{Name: "개발자동차", Position: 0}

	car.Move()

	assert.Equal(
		t,
		1,
		car.Position,
		"Expected car position to be %d, but got %d", car.Position+1, car.Position,
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

type AlwaysTrueRule struct{}

func (r AlwaysTrueRule) Check(n int) bool {
	return true
}

func TestPlayGame(t *testing.T) {
	alwaysTrueRule := AlwaysTrueRule{}

	car1 := &Car{Name: "Car1", Position: 0}
	car2 := &Car{Name: "Car2", Position: 0}
	game := NewCarGame([]*Car{car1, car2}, alwaysTrueRule)
	game.PlayGame(5)

	for _, car := range game.Cars {
		assert.Equal(
			t,
			car.Position,
			1,
			"Expected car position to be %d, but got %d", 1, car.Position,
		)
	}
}
