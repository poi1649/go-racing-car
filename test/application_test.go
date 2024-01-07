package test

import (
	"github.com/poi1649/go-racing-car/racingcar"
	"github.com/stretchr/testify/assert"

	"testing"
)

// 예시 테스트 함수, 구현 시 제거
func TestHello(t *testing.T) {
	expected := "Hello, Racing Car!"
	actual := racingcar.Hello()
	assert.Equal(
		t,
		expected,
		actual,
		"they should be equal",
	)
}
