package test

import (
	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCars(t *testing.T) {
	car1, _ := racingcar.NewCar("test1")
	car2, _ := racingcar.NewCar("test2")
	car3, _ := racingcar.NewCar("test3")
	cars := racingcar.NewCars([]*racingcar.Car{car1, car2, car3})
	assert.Equal(t, 3, len(cars.Cars))
}
