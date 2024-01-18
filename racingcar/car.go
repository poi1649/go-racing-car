package racingcar

import (
	"errors"
	"unicode/utf8"
)

type Car struct {
	name     string
	position int
}

func NewCar(name string) (*Car, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	return &Car{name, 0}, nil
}

func validateName(name string) error {
	nameLength := utf8.RuneCountInString(name)
	switch {
	case nameLength == 0:
		return errors.New("empty name")
	case nameLength > 5:
		return errors.New("too long name")
	}
	return nil
}

func (car *Car) MoveForward() {
	car.position += 1
}

func (car Car) Name() string {
	return car.name
}

func (car Car) Position() int {
	return car.position
}
