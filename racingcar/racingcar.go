package racingcar

import (
	"math/rand"
	"time"
)

type Car struct {
	Name     string
	Position int
}

func (c *Car) Move() {
	c.Position += 1
}

func (c Car) Stop() Car {
	return c
}

type CarGame struct {
	Cars []*Car
	rule Rule
}

func NewCarGame(cars []*Car, rule Rule) CarGame {
	return CarGame{Cars: cars, rule: rule}
}

// Rule 인터페이스
type Rule interface {
	Check(n int) bool
}

// RandomRule은 Rule 인터페이스를 구현하는 구조체
type RandomRule struct{}

func (r RandomRule) Check(n int) bool {
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(10)
	return randomValue >= n
}

// PlayGame 메서드
func (cg *CarGame) PlayGame(n int) {
	for i := range cg.Cars {
		if cg.rule.Check(n) {
			cg.Cars[i].Move()
		}
	}
}
