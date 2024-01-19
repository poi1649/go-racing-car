package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/poi1649/go-racing-car/racingcar"
)

type Cli struct {
	reader *bufio.Reader
}

func NewCli(reader *bufio.Reader) *Cli {
	return &Cli{reader: reader}
}

func (cli *Cli) GetCarNames() []string {
	input, err := cli.reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.Split(input, ",")
}

func (cli *Cli) GetTryCount() int {
	input, err := cli.reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	tryCount, atoiErr := strconv.Atoi(strings.TrimSpace(input))
	if atoiErr != nil {
		panic("시도할 회수는 숫자만 가능합니다.")
	}

	return tryCount
}

func (cli *Cli) PrintCurrentStatus(game *racingcar.Game) {
	for _, car := range game.Cars {
		fmt.Printf("%s : %s\n", car.Name(), strings.Repeat("-", car.Position()))
	}
	fmt.Println()
}
