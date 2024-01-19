package main

import (
	"bufio"
	"fmt"
	"github.com/thedevluffy/go-racing-car/racingcar"
	"os"
	"strings"
)

func main() {
	// 사용자 입력 처리
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("경주 할 자동차 이름(이름은 쉼표(,) 기준으로 구분):")
	scanner.Scan()
	carNames := strings.Split(scanner.Text(), ",")

	fmt.Println("시도할 회수:")
	scanner.Scan()
	var attempts int
	fmt.Sscanf(scanner.Text(), "%d", &attempts)

	// CarGame 초기화
	cars := make([]*racingcar.Car, len(carNames))
	for i, name := range carNames {
		cars[i] = &racingcar.Car{Name: strings.TrimSpace(name)}
	}
	randomRule := racingcar.RandomRule{}
	game := racingcar.NewCarGame(cars, randomRule)

	// 게임 실행 및 결과 출력
	for i := 0; i < attempts; i++ {
		game.PlayGame(4)
		fmt.Printf("시도 #%d:\n", i+1)
		for _, car := range game.Cars {
			fmt.Printf("%s : %s\n", car.Name, car)
		}
	}

	// 우승자 찾기
	var maxPosition int
	for _, car := range game.Cars {
		if car.Position > maxPosition {
			maxPosition = car.Position
		}
	}
	var winners []string
	for _, car := range game.Cars {
		if car.Position == maxPosition {
			winners = append(winners, car.Name)
		}
	}

	// 우승자 출력
	if len(winners) == 1 {
		fmt.Println("최종 우승자 :", winners[0])
	} else {
		fmt.Println("최종 우승자 :", strings.Join(winners, ", "))
	}
}
