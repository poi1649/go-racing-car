package racingcar

import (
	"errors"
	"math/rand"
)

type FuelGenerator interface {
	Fuel() (int, error)
}

type RandomFuelGenerator struct {
}

func NewRandomFuelGenerator() *RandomFuelGenerator {
	return &RandomFuelGenerator{}
}

func (r *RandomFuelGenerator) Fuel() (int, error) {
	return rand.Intn(10), nil
}

type FixedFuelGenerator struct {
	FuelValue []int
	index     int
}

func NewFixedFuelGenerator(fuelValue []int) *FixedFuelGenerator {
	return &FixedFuelGenerator{FuelValue: fuelValue, index: 0}
}

func (f *FixedFuelGenerator) Fuel() (int, error) {
	if f.index >= len(f.FuelValue) {
		return 0, errors.New("제공된 연료 길이보다 많은 연료를 요청했습니다")
	}
	fuel := f.FuelValue[f.index]
	f.index++
	return fuel, nil
}
