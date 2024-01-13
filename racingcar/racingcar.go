package racingcar

type Car struct {
	Name     string
	Position int
}

func (c Car) Move() Car {
	return Car{Name: c.Name, Position: c.Position + 1}
}

func (c Car) Stop() Car {
	return c
}
