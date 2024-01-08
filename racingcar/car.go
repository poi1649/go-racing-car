package racingcar

import "errors"

type Car struct {
	Name     string
	Position int
}

const CarFuelThreshold = 4
const CarMoveStep = 1
const CarNameMinLength = 0
const CarNameMaxLength = 5
const CarMaxFuel = 9
const CarMinFuel = 0

func NewCar(name string) (*Car, error) {
	if !checkValidCarName(name) {
		return nil, errors.New("차 이름은 1자 이상 5자 이하만 가능합니다")
	}
	return &Car{Name: name, Position: 0}, nil
}

func checkValidCarName(name string) bool {
	return CarNameMinLength < len(name) && len(name) <= CarNameMaxLength
}

func (c *Car) Move(fuel int) error {
	if !checkValidFuel(fuel) {
		return errors.New("연료는 0 이상 9 이하만 가능합니다")
	}
	if fuel >= CarFuelThreshold {
		c.Position += CarMoveStep
		return nil
	}
	return nil
}

func checkValidFuel(fuel int) bool {
	return CarMinFuel <= fuel && fuel <= CarMaxFuel
}
