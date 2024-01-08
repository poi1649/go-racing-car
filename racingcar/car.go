package racingcar

import "errors"

type Car struct {
	Name     string
	Position int
}

func NewCar(name string) (*Car, error) {
	if !checkValidCarName(name) {
		return nil, errors.New("차 이름은 1자 이상 5자 이하만 가능합니다")
	}
	return &Car{Name: name, Position: 0}, nil
}

func checkValidCarName(name string) bool {
	return 0 < len(name) && len(name) <= 5
}