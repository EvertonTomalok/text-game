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

	for n := 0; n < len(testCase.RandomNumber); n++ {
		question := <-results
		questionsContainer.AddQuestion(question)
	}

	if len(questionsContainer.Questions) != testCase.ExpectQuestionsNumber {
		t.Errorf(
			"the expected number of questions were %d and received %d",
			len(questionsContainer.Questions),
			testCase.ExpectQuestionsNumber,
		)
	}
}
