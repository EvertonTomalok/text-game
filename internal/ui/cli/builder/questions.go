package builder

import (
	"github.com/evertotomalok/text-game/internal/app/shared"
	"github.com/evertotomalok/text-game/internal/core/domain"
	"github.com/evertotomalok/text-game/internal/core/errs"
	"github.com/evertotomalok/text-game/pkg/utils"
)

func CreateQuestions(rounds uint) ([]domain.Question, error) {
	if rounds < 15 || rounds > 30 {
		return make([]domain.Question, 0), new(errs.InvalidRoundsNumber)
	}

	/*
	* First of all, we populate the question with minimum rounds value acceptable, in this case will be 8 (all questions available)
	* Each question will take 2 rounds:
	* 	First: The main QUESTION
	*	Second: The complementary QUESTION
	*
	* This way we'll have at least 16 question.
	 */
	var questions []domain.Question = suffleQuestions()

	if !utils.IntegerIsEven(int64(rounds)) {
		rounds++
	}

	// Minimum rounds = 15
	// If it's odd, we add 1 to number of rounds to retrieve an integer rounds number
	baseQuestions := 8
	necessaryQuestion := rounds/2 - uint(baseQuestions)
	var additionalQuestions []domain.Question = suffleQuestions()

	/* Now populate the rest of the questions that the porgram will need, to achieve from 15 to 30 available rounds */
	for i := 0; i < int(necessaryQuestion); i++ {
		questions = append(questions, additionalQuestions[i])
	}

	return questions, nil
}

func getShuffledQuestionsNumbers() []uint {
	numbers := make([]uint, len(shared.MainQuestions))

	for i := 1; i < len(shared.MainQuestions); i++ {
		numbers[i] = uint(i)
	}

	return utils.Shuffle(numbers)
}

func suffleQuestions() []domain.Question {
	workersNum := 5
	var questions []domain.Question

	randomNumbers := getShuffledQuestionsNumbers()
	numTasks := len(randomNumbers)

	tasks, results := CreateChannels(numTasks)

	CreateWorkers(tasks, results, workersNum)
	PopulateTasks(tasks, randomNumbers)

	// Listen to the results channel to populate the questions slice
	for m := 0; m < numTasks; m++ {
		question := <-results
		questions = append(questions, question)
	}

	return questions
}
