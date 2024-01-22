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
	carNames := cli.MustGetCarNames()

	fmt.Println("시도할 회수는 몇회인가요?")
	tryCount := cli.MustGetTryCount()

	game, err := racingcar.NewGame(carNames, tryCount)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n실행 결과\n")
	for !game.IsEnd() {
		game.PlayTurn(racingcar.DefalutMoveStrategy)
		cli.PrintCurrentStatus(game)
	}
	cli.PrintWinners(game)
}
