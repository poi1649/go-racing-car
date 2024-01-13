package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thedevluffy/go-racing-car/racingcar"

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
