package model

import (
	"github.com/poi1649/go-racing-car/racingcar/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCar(t *testing.T) {
	car, err := model.NewCar("test")
	assert.Equal(t, nil, err)
	assert.Equal(t, "test", car.Name)
}

func TestCreateCarWithBlankName(t *testing.T) {
	_, err := model.NewCar("")
	errMsg := "차 이름은 1자 이상 5자 이하만 가능합니다"
	assert.Equal(t, errMsg, err.Error())
}

func TestCreateCarWithOverLengthName(t *testing.T) {
	_, err := model.NewCar("123456")
	errMsg := "차 이름은 1자 이상 5자 이하만 가능합니다"
	assert.Equal(t, errMsg, err.Error())
}

func TestCarPosition(t *testing.T) {
	car, _ := model.NewCar("test")
	assert.Equal(t, 0, car.Position)
}

func TestCarMove(t *testing.T) {
	car, _ := model.NewCar("test")
	err := car.Move(6)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, car.Position)
}

func TestCarNotMove(t *testing.T) {
	car, _ := model.NewCar("test")
	err := car.Move(3)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, car.Position)
}

func TestCarMoveWithInvalidFuel(t *testing.T) {
	car, _ := model.NewCar("test")
	err := car.Move(10)
	errMsg := "연료는 0 이상 9 이하만 가능합니다"
	assert.Equal(t, errMsg, err.Error())
	assert.Equal(t, 0, car.Position)
}
