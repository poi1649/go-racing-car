package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/poi1649/go-racing-car/cmd"
	"github.com/poi1649/go-racing-car/racingcar"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
		}
	}()

	cli := cmd.NewCli(bufio.NewReader(os.Stdin))

	fmt.Println("경주할 자동차 이름을 입력하세요.(이름은 쉼표(,) 기준으로 구분)")
	carNames := cli.GetCarNames()

	game, err := racingcar.NewGame(carNames)
	if err != nil {
		panic(err)
	}

	fmt.Println("시도할 회수는 몇회인가요?")
	tryCount := cli.GetTryCount()

	if !validateGameCount(tryCount) {
		panic("시도할 회수는 1이상 100이하의 자연수만 가능합니다.")
	}

	fmt.Printf("\n실행 결과\n")
	for turn := 0; turn < tryCount; turn++ {
		game.PlayTurn(racingcar.DefalutMoveStrategy)
		cli.PrintCurrentStatus(game)
	}
}

func validateGameCount(count int) bool {
	return count > 0 && count <= 100
}
