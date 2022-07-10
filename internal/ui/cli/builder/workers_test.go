package builder_test

import (
	"testing"

	. "github.com/evertotomalok/text-game/internal/ui/cli/builder"
)

type TestCase struct {
	QuestionsNumber       int
	ExpectQuestionsNumber int
}

func TestMakeChannels(t *testing.T) {
	channelNum := 10
	tasks, results := CreateChannels(channelNum)

	if cap(tasks) != channelNum {
		t.Errorf("the capability of the channel should be %d and is %d", channelNum, cap(tasks))
	}

	if cap(results) != channelNum {
		t.Errorf("the capability of the channel should be %d and is %d", channelNum, cap(results))
	}
}

func TestWorkerPoolInitiate(t *testing.T) {

}
