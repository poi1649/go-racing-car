package racingcar

import (
	"fmt"
	"math/rand"
	"strings"
)

type CarRace struct {
	CarNames []string
	CarsInfo map[string]int
	Count    int
}

func NewCarRace(carNames []string, count int) *CarRace {
	carsInfo := make(map[string]int)
	for _, name := range carNames {
		carsInfo[strings.TrimSpace(name)] = 0
	}

	return &CarRace{
		CarNames: carNames,
		CarsInfo: carsInfo,
		Count:    count,
	}
}

func (c *CarRace) Start() []string {
	fmt.Println("실행 결과")
	for i := 0; i < c.Count; i++ {
		for carName := range c.CarsInfo {
			if rand.Float64() < 0.5 {
				c.CarsInfo[carName]++
			}
			fmt.Println(carName, " : ", strings.Repeat("-", c.CarsInfo[carName]))
		}
		fmt.Println()
	}

	return c.findWinners()
}

func (c *CarRace) findWinners() []string {
	maxValue := 0
	var winners []string

	for carName, value := range c.CarsInfo {
		if value > maxValue {
			maxValue = value
			winners = []string{carName}
		} else if value == maxValue {
			winners = append(winners, carName)
		}
	}

	return winners
}

func PrintWinners(winners []string) {
	fmt.Print("최종 우승자: ")
	for i, winner := range winners {
		if i == len(winners)-1 {
			fmt.Print(winner)
		} else {
			fmt.Print(winner, ", ")
		}
	}
	fmt.Println()
}
