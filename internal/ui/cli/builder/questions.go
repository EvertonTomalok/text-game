package builder

import (
	"sync"

	"github.com/evertotomalok/text-game/internal/app/shared"
	"github.com/evertotomalok/text-game/internal/core/domain"
	"github.com/evertotomalok/text-game/internal/core/errs"
	"github.com/evertotomalok/text-game/pkg/utils"
)

type QuestionsContainer struct {
	mu        sync.Mutex
	Questions []domain.Question
}

func (q *QuestionsContainer) AddQuestion(question domain.Question) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Questions = append(q.Questions, question)
}

func CreateQuestions(rounds uint) ([]domain.Question, error) {
	if rounds < 15 || rounds > 30 {
		return make([]domain.Question, 0), new(errs.InvalidRoundsNumber)
	}

	/*
	* First of all, we'll populate the question with the minimum rounds value acceptable, in this case will be 8 (all questions available)
	* Each question will take 2 rounds:
	* 	First: The main QUESTION
	*	Second: The complementary QUESTION
	*
	* On this way we'll have at least 16 question.
	 */

	questions, err := suffleQuestions(
		len(shared.MainQuestions),
	)

	if err != nil {
		return make([]domain.Question, 0), err
	}

	if !utils.IntegerIsEven(int64(rounds)) {
		rounds++
	}

	// Minimum rounds = 15
	// If it's odd, we add 1 to number of rounds to retrieve an integer rounds number
	if rounds > 16 {
		baseQuestions := uint(8)
		necessaryQuestion := rounds/2 - baseQuestions

		additionalQuestions, err := suffleQuestions(
			int(necessaryQuestion),
		)

		if err != nil {
			return make([]domain.Question, 0), err
		}

		/* Now populate the rest of the questions that the porgram will need, to achieve from 15 to 30 available rounds */
		questions = append(questions, additionalQuestions...)
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

func suffleQuestions(numQuestions int) ([]domain.Question, error) {

	if numQuestions > len(shared.MainQuestions) {
		numQuestions = len(shared.MainQuestions)
	}

	randomNumbers := getShuffledQuestionsNumbers()[:numQuestions]

	results := QuestionsWorkerPoolInitiate(numQuestions, randomNumbers)

	questionsContainer := &QuestionsContainer{}

	// Listen to the results channel to populate the questions' slice
	for m := 0; m < numQuestions; m++ {
		question := <-results

		if question.ID == 0 {
			return questionsContainer.Questions, new(errs.InvalidQuestion)
		}

		questionsContainer.AddQuestion(question)
	}

	return questionsContainer.Questions, nil
}
