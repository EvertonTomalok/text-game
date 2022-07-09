package builder_test

import (
	"testing"

	. "github.com/evertotomalok/text-game/internal/ui/cli/builder"
)

type ScenarioTest struct {
	RoundNumber       uint
	ExpectedQuestions uint
}

// Expecting to be ok
func TestQuestionsBuilder_OK(t *testing.T) {
	scenariosTest := []ScenarioTest{
		{
			RoundNumber:       15,
			ExpectedQuestions: 8,
		},
		{
			RoundNumber:       16,
			ExpectedQuestions: 8,
		},
		{
			RoundNumber:       17,
			ExpectedQuestions: 9,
		},
		{
			RoundNumber:       20,
			ExpectedQuestions: 10,
		},
		{
			RoundNumber:       25,
			ExpectedQuestions: 13,
		},
		{
			RoundNumber:       30,
			ExpectedQuestions: 15,
		},
	}

	for _, scenarioTest := range scenariosTest {
		questions, _ := CreateQuestions(uint(scenarioTest.RoundNumber))
		if len(questions) != int(scenarioTest.ExpectedQuestions) {
			t.Errorf("the number of Expected Questions was %d but received %d", scenarioTest.ExpectedQuestions, len(questions))
		}
	}
}

// Expecting to throw error building questions
func TestQuestionsBuilder_NOT_OK(t *testing.T) {
	scenariosTest := []ScenarioTest{
		{
			RoundNumber: 10,
		},
		{
			RoundNumber: 0,
		},
		{
			RoundNumber: 8,
		},
		{
			RoundNumber: 31,
		},
		{
			RoundNumber: 40,
		},
	}

	for _, scenarioTest := range scenariosTest {
		_, err := CreateQuestions(uint(scenarioTest.RoundNumber))

		if err == nil {
			t.Error("it must receive an error, but it didn't receive")
		}
	}
}
