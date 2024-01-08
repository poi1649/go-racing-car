package test

import (
	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCar(t *testing.T) {
	car, err := racingcar.NewCar("test")
	assert.Equal(t, nil, err)
	assert.Equal(t, "test", car.Name)
}

func TestCreateCarWithBlankName(t *testing.T) {
	_, err := racingcar.NewCar("")
	errMsg := "차 이름은 1자 이상 5자 이하만 가능합니다"
	assert.Equal(t, errMsg, err.Error())
}

func TestCreateCarWithOverLengthName(t *testing.T) {
	_, err := racingcar.NewCar("123456")
	errMsg := "차 이름은 1자 이상 5자 이하만 가능합니다"
	assert.Equal(t, errMsg, err.Error())
}

func TestCarPosition(t *testing.T) {
	car, _ := racingcar.NewCar("test")
	assert.Equal(t, 0, car.Position)
}

func TestCarMove(t *testing.T) {
	car, _ := racingcar.NewCar("test")
	car.Move()
	assert.Equal(t, 1, car.Position)
}
