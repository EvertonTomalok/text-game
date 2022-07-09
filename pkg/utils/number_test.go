package utils_test

import (
	"testing"

	. "github.com/evertotomalok/text-game/pkg/utils"
)

func TestIsEven_True(t *testing.T) {
	if ok := IntegerIsEven(int64(10)); !ok {
		t.Fail()
	}
}

func TestIsEven_False(t *testing.T) {
	if ok := IntegerIsEven(int64(11)); ok {
		t.Fail()
	}
}
