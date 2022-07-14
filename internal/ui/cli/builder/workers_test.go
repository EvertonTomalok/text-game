package builder_test

import (
	"testing"

	. "github.com/evertotomalok/text-game/internal/ui/cli/builder"
)

type TestCase struct {
	RandomNumber          []uint // 0 <-> 7
	ExpectQuestionsNumber int
}

func TestMakeChannels(t *testing.T) {
	channelNum := 10
	tasks, results := CreateChannels(channelNum)
	defer close(tasks)
	defer close(results)

	if cap(tasks) != channelNum {
		t.Errorf("the capability of the channel should be %d and is %d", channelNum, cap(tasks))
	}

	if cap(results) != channelNum {
		t.Errorf("the capability of the channel should be %d and is %d", channelNum, cap(results))
	}
}

func TestQuestionsWorkerPoolInitiate(t *testing.T) {

	testCase := TestCase{
		RandomNumber:          []uint{5, 3, 4},
		ExpectQuestionsNumber: 3,
	}

	questionsContainer := &QuestionsContainer{}
	results := QuestionsWorkerPoolInitiate(len(testCase.RandomNumber), testCase.RandomNumber)

	questions, err := ListenResults(len(testCase.RandomNumber), results)

	if err != nil {
		t.Errorf(
			"some error ocurred %+v",
			err,
		)
	}

	if len(questions) != testCase.ExpectQuestionsNumber {
		t.Errorf(
			"the expected number of questions were %d and received %d",
			len(questionsContainer.Questions),
			testCase.ExpectQuestionsNumber,
		)
	}
}

func TestMakeQuestionsUsingWorkerPool_MAX_QUESTIONS_OK(t *testing.T) {
	lenQuestions := 8
	questions, err := MakeQuestionsUsingWorkerPool(lenQuestions)

	if err != nil {
		t.Errorf("some error ocurred %+v", err)
	}

	if lenQuestions != len(questions) {
		t.Errorf("it was expected %d but received %d", lenQuestions, len(questions))
	}
}

func TestMakeQuestionsUsingWorkerPool_MAX_QUESTIONS_EIGHT(t *testing.T) {
	expectedQuestions := 8

	lenQuestions := 10
	questions, err := MakeQuestionsUsingWorkerPool(lenQuestions)

	if err != nil {
		t.Errorf("some error ocurred %+v", err)
	}

	/* Even we pass more than 8 as lenght of question, we'll only produce 8 questions */
	if expectedQuestions != len(questions) {
		t.Errorf("it was expected %d but received %d", lenQuestions, len(questions))
	}
}
