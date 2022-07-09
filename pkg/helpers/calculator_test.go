package helpers_test

import (
	"testing"

	"github.com/evertotomalok/text-game/internal/app/config"
	. "github.com/evertotomalok/text-game/pkg/helpers"
)

func TestDecreaseCalculator(t *testing.T) {
	flow := config.Flow{
		Rounds:          16,
		DebugComplexity: 1,
	}
	DecrementOneToComplexity(&flow, 1)

	if flow.DebugComplexity != 0 {
		t.Fail()
	}
}

func TestIncreaseCalculator(t *testing.T) {
	flow := config.Flow{
		Rounds:          16,
		DebugComplexity: -1,
	}
	IncrementOneToComplexity(&flow, 1)

	if flow.DebugComplexity != 0 {
		t.Fail()
	}
}
