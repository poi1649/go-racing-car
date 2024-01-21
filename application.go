package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("경주할 자동차 이름을 입력하세요.(이름은 쉼표(,) 기준으로 구분")
	scanner.Scan()
	carNames := strings.Split(scanner.Text(), ",")
	carsInfo := make(map[string]int)
	for _, name := range carNames {
		carsInfo[strings.TrimSpace(name)] = 0
	}

	fmt.Println("시도할 회수는 몇회인가요? ")
	scanner.Scan()
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("입력한 회수를 정수로 변환하는데 실패했습니다.")
		return
	}

	fmt.Println("실행 결과")
	for i := 0; i < count; i++ {
		for carName := range carsInfo {
			if rand.Float64() < 0.5 {
				carsInfo[carName]++
			}
			fmt.Println(carName, " : ", strings.Repeat("-", carsInfo[carName]))
		}
		fmt.Println()
	}

	maxValue := 0
	var winners []string

	for carName, value := range carsInfo {
		if value > maxValue {
			maxValue = value
			winners = []string{carName}
		} else if value == maxValue {
			winners = append(winners, carName)
		}
	}

	// Print all keys with the maximum value
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
