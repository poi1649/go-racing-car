package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/poi1649/go-racing-car/racingcar"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("경주할 자동차 이름을 입력하세요.(이름은 쉼표(,) 기준으로 구분)")
	carNameInput, _ := reader.ReadString('\n')

	carNames := strings.Split(carNameInput, ",")
	game, err := racingcar.NewGame(carNames)
	if err != nil {
		panic(err)
	}

	fmt.Println("시도할 회수는 몇회인가요?")
	tryCountInput, _ := reader.ReadString('\n')
	tryCount, atoiErr := strconv.Atoi(strings.TrimSpace(tryCountInput))
	if atoiErr != nil {
		panic("시도할 회수는 숫자만 가능합니다.")
	}
	if !validateGameCount(tryCount) {
		panic("시도할 회수는 1이상 100이하의 자연수만 가능합니다.")
	}

	for turn := 0; turn < tryCount; turn++ {
		game.PlayTurn(racingcar.DefalutMoveStrategy)
		for _, car := range game.Cars {
			printCar(car)
		}
		fmt.Println()
	}
}

func validateGameCount(count int) bool {
	return count > 0 && count <= 100
}

func printCar(car racingcar.Car) {
	fmt.Printf("%s : %s\n", car.Name(), strings.Repeat("-", car.Position()))
}
