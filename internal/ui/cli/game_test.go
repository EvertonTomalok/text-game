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
	FlowControllerAdapter = new(cliMocker.FlowValidatorMockError)

	mockGetInput := new(mocks.InputInterface)
	mockGetInput.On("GetValidInput").Return(1)
	GetInputUserAdapter = mockGetInput
	_ = StartGame(&flow)
}
