package racingcar

import "math/rand"

const (
	defaultMoveThreshod = 4
	defaultMoveNumRange = 10
)

type MoveStrategy func() bool

func DefalutMoveStrategy() bool {
	return rand.Intn(defaultMoveNumRange) >= defaultMoveThreshod
}

// 위의 함수를 바는 방법과 아래의 인터페이스와 빈 struct를 사용하는 방법중 고민
// 위의 함수를 이용한 방법은 좀 더 간단하고 단순하게 사용할 수 있다는 장점이 있는것고
// 아래의 인터페이스를 이용한 방법은 좀 더 명시적이고, 확장성이 있는것 같습니다.
// 이동에 대한 더 복잡한 전략이 생길 여지가 있다면 아래의 방법이 좋을 것 같지만 일단 간단하게 미션진행을 위해 위의 방법으로 구현해보았습니다.
// 의견이 궁금해서 주석처리하고 남겨두었습니다. 의견 있으면 공유하고싶어요!!
/*
type MoveStrategy interface {
	IsMovable() bool
}

type DefaultStrategy struct{}

func (DefaultStrategy) IsMovable() bool {
	return rand.Intn(10) >= defaultMoveThresho
}

// function at game.go
func playTurn(car *Car, moveStrategy MoveStrategy) {
	if moveStrategy.IsMovable() {
		car.MoveForward()
	}
}
*/
