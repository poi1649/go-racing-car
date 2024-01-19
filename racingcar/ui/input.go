package ui

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Func func() (interface{}, error)

func GetNamesFromReader() (interface{}, error) {
	println("경주할 자동차 이름을 입력하세요(이름은 쉼표(,) 기준으로 구분).")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return nil, err
	}
	names := strings.Split(input, ",")
	return names, validateNames(names)
}

func validateNames(names []string) error {
	if len(names) < 2 {
		return errors.New("2개 이상의 이름을 입력해주세요")
	}
	if hasDuplicates(names) {
		return errors.New("중복된 이름을 입력할 수 없습니다")
	}
	return nil
}

func hasDuplicates(elements []string) bool {
	seen := make(map[string]bool)
	for _, element := range elements {
		if seen[element] {
			return true
		}
		seen[element] = true
	}
	return false
}

func GetTrialFromReader() (interface{}, error) {
	println("시도할 회수는 몇회인가요?")
	var input string
	_, err := fmt.Scanln(&input)
	parseInt, err := strconv.Atoi(input)
	return parseInt, validateTrial(parseInt, err)
}

func validateTrial(round int, err error) error {
	if err != nil {
		return err
	}
	if round < 1 {
		return errors.New("1 이상의 숫자를 입력해주세요")
	}
	return nil
}
