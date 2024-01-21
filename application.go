package main

import (
	"bufio"
	"fmt"
	"github.com/poi1649/go-racing-car/racingcar"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("경주할 자동차 이름을 입력하세요.(이름은 쉼표(,) 기준으로 구분")
	scanner.Scan()
	carNames := strings.Split(scanner.Text(), ",")

	fmt.Println("시도할 회수는 몇회인가요? ")
	scanner.Scan()
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("입력한 회수를 정수로 변환하는데 실패했습니다.")
		return
	}

	race := racingcar.NewCarRace(carNames, count)
	winners := race.Start()
	racingcar.PrintWinners(winners)
}
