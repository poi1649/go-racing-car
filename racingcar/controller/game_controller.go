package controller

import (
	"github.com/poi1649/go-racing-car/racingcar/model"
	"github.com/poi1649/go-racing-car/racingcar/ui"
)

func StartGame() {
	names := handleIOError(ui.GetNamesFromReader)
	trial := handleIOError(ui.GetTrialFromReader)
	game := initializeGame(names.([]string))
	for i := 0; i < trial.(int); i++ {
		result, err := game.Run()
		panicIfError(err)
		ui.PrintRunResultOutput(result)
	}
	winnersName := game.GetWinnersName()
	ui.PrintWinnerOutput(winnersName)
}

func initializeGame(names []string) *model.Game {
	game, err := model.NewGame(names, model.NewRandomFuelGenerator())
	panicIfError(err)
	return game
}

func handleIOError(uiFunction ui.Func) interface{} {
	value, err := uiFunction()
	panicIfError(err)
	return value
}

func panicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
