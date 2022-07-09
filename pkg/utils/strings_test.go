package utils_test

import (
	"testing"

	. "github.com/evertotomalok/text-game/pkg/utils"
)

func TestSeparator(t *testing.T) {
	separator := Separator("-", 50)

	if len(separator) != 50 {
		t.Fail()
	}
}
