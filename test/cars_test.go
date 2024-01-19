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

func TestGenerateCarsWithMethod(t *testing.T) {
	cars, _ := racingcar.GenerateCars([]string{"test1", "test2", "test3"})
	assert.Equal(t, 3, len(cars.Cars))
}

func TestGenerateCarsWithInvalidName(t *testing.T) {
	_, err := racingcar.GenerateCars([]string{"test1", "test2", ""})
	errMsg := "차 이름은 1자 이상 5자 이하만 가능합니다"
	assert.Equal(t, errMsg, err.Error())
}

func TestGetWinners(t *testing.T) {
	car1, _ := racingcar.NewCar("test1")
	car2, _ := racingcar.NewCar("test2")
	car3, _ := racingcar.NewCar("test3")
	cars := racingcar.NewCars([]*racingcar.Car{car1, car2, car3})
	car1.Move(3)
	car2.Move(5)
	car3.Move(1)
	winners := cars.GetWinners()
	assert.Equal(t, 1, len(winners))
	assert.Equal(t, "test2", winners[0].Name)
}

func TestGetWinnersWithSamePosition(t *testing.T) {
	car1, _ := racingcar.NewCar("test1")
	car2, _ := racingcar.NewCar("test2")
	car3, _ := racingcar.NewCar("test3")
	cars := racingcar.NewCars([]*racingcar.Car{car1, car2, car3})
	car1.Move(3)
	car2.Move(5)
	car3.Move(5)
	winners := cars.GetWinners()
	assert.Equal(t, 2, len(winners))
	assert.Equal(t, "test2", winners[0].Name)
	assert.Equal(t, "test3", winners[1].Name)
}
