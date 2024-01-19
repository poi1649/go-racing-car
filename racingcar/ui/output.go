package ui

import (
	"github.com/poi1649/go-racing-car/racingcar/model"
	"strings"
)

func PrintResultMessage() {
	println("실행 결과")
}

func PrintRunResultOutput(runResult *model.RunResult) {
	output := ""
	for carName, position := range runResult.CarNameToPosition {
		output += carName + " : "
		for i := 0; i < position; i++ {
			output += "-"
		}
		output += "\n"
	}
	println(output)
}

func PrintWinnerOutput(winnersName []string) {
	output := "최종 우승자 : "
	names := strings.Join(winnersName, ", ")
	output += names
	println(output)
}

func PrintErrorMessage(err error) {
	println(err.Error())
}
