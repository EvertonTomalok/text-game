package cli_test

import (
	"testing"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/internal/app/shared"
	"github.com/evertotomalok/text-game/internal/core/ports/mocks"
	. "github.com/evertotomalok/text-game/internal/ui/cli"
	cliMocker "github.com/evertotomalok/text-game/internal/ui/cli/adapters/mocks"
)

type TestMainQuestionCase struct {
	IndexMainQuestion       uint
	ExpectedComplementaryId int
}

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

	mockGetInput := new(mocks.InputGateway)
	mockGetInput.On("GetValidInput").Return(uint(1)) // It mocks to return every time the FirstAlternative
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

func TestProcessQuestion(t *testing.T) {
	mockGetInput := new(mocks.InputGateway)
	mockGetInput.On("GetValidInput").Return(uint(1)) // It mocks to return every time the FirstAlternative
	GetInputUserAdapter = mockGetInput

	flow := config.Flow{
		Rounds:          15,
		DebugComplexity: 0,
	}

	testCases := []TestMainQuestionCase{
		{
			IndexMainQuestion:       0, // key of map in shared.MainQuestions
			ExpectedComplementaryId: 1, // key of map in shared.ComplementaryQuestions.
			// First MainQuestion links to the ComplementaryQuestions of key 1 and 2
		},
		{
			IndexMainQuestion:       1,
			ExpectedComplementaryId: 3,
		},
		{
			IndexMainQuestion:       2,
			ExpectedComplementaryId: 5,
		},
		{
			IndexMainQuestion:       3,
			ExpectedComplementaryId: 7,
		},
		{
			IndexMainQuestion:       4,
			ExpectedComplementaryId: 9,
		},
		{
			IndexMainQuestion:       5,
			ExpectedComplementaryId: 11,
		},
		{
			IndexMainQuestion:       6,
			ExpectedComplementaryId: 13,
		},
		{
			IndexMainQuestion:       7,
			ExpectedComplementaryId: 15,
		},
	}

	for _, testCase := range testCases {
		question := shared.MainQuestions[testCase.IndexMainQuestion]
		complementaryQuestion, err := ProcessQuestion(question, &flow)

		if err != nil {
			t.Error("it must not return error")
		}

		if complementaryQuestion.ID != testCase.ExpectedComplementaryId {
			t.Errorf(
				"the ComplementaryQuestion.ID expected was %d but received %d",
				testCase.ExpectedComplementaryId,
				complementaryQuestion.ID,
			)
		}
	}
}

func TestProcessComplementaryQuestion(t *testing.T) {
	mockGetInput := new(mocks.InputGateway)
	mockGetInput.On("GetValidInput").Return(uint(1)) // It mocks to return every time the FirstAlternative
	GetInputUserAdapter = mockGetInput

	flow := config.Flow{
		Rounds:          15,
		DebugComplexity: 0,
	}

	complementaryQuestion := shared.ComplementaryQuestions[1]

	if err := ProcessComplementaryQuestion(complementaryQuestion, &flow); err != nil {
		t.Error("this function couldn't return error")
	}
}
