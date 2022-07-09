package cli_test

import (
	"testing"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/internal/core/ports/mocks"
	. "github.com/evertotomalok/text-game/internal/ui/cli"
	cliMocker "github.com/evertotomalok/text-game/internal/ui/cli/adapters/mocks"
)

func TestGame_Flow_Validator_Error(t *testing.T) {
	FlowControllerAdapter = new(cliMocker.FlowValidatorMockError)
	flow := config.Flow{
		Rounds: 14,
	}
	if err := StartGame(&flow); err == nil {
		t.Error("It should return a ERROR.")
	}
}

func TestGame_OK(t *testing.T) {
	flow := config.Flow{
		Rounds: 15,
	}

	FlowControllerAdapter = new(cliMocker.FlowValidatorMock)

	mockGetInput := new(mocks.InputInterface)
	mockGetInput.On("GetValidInput").Return(uint(1))
	GetInputUserAdapter = mockGetInput

	if err := StartGame(&flow); err != nil {
		t.Fail()
	}
}

func TestComputeDebugComplexity_INCREASE(t *testing.T) {
	flow := config.Flow{
		Rounds:          15,
		DebugComplexity: 0,
	}
	ComputeDebugComplexity(&flow, 1)

	if flow.DebugComplexity != 1 {
		t.Fail()
	}
}

func TestComputeDebugComplexity_DECREASE(t *testing.T) {
	flow := config.Flow{
		Rounds:          15,
		DebugComplexity: 0,
	}
	ComputeDebugComplexity(&flow, -1)

	if flow.DebugComplexity != -1 {
		t.Fail()
	}
}
