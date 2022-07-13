package helpers_test

import (
	"testing"

	"github.com/evertotomalok/text-game/internal/application/config"
	. "github.com/evertotomalok/text-game/pkg/helpers"
)

func TestFlowValidator_OK(t *testing.T) {
	flow := config.Flow{
		Rounds:          16,
		DebugComplexity: 0,
	}
	err := FlowValidator(&flow)
	if err != nil {
		t.Fail()
	}
}
