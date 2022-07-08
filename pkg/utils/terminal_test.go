package utils_test

import (
	"testing"

	. "github.com/evertotomalok/text-game/pkg/utils"
)

func TestClearTerminal_OK(t *testing.T) {
	ClearTerminal()
}
