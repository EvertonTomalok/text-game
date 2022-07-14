package builder

import (
	"sync"

	"github.com/evertotomalok/text-game/internal/core/domain"
	"github.com/evertotomalok/text-game/internal/core/errs"
	"github.com/evertotomalok/text-game/pkg/utils"
)

type QuestionsContainer struct {
	mu        sync.Mutex
	Questions []domain.Question
}

func (q *QuestionsContainer) AddQuestion(question ...domain.Question) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Questions = append(q.Questions, question...)
}

func CreateQuestions(rounds uint) ([]domain.Question, error) {
	if rounds < 15 || rounds > 30 {
		return make([]domain.Question, 0), new(errs.InvalidRoundsNumber)
	}

	/*
	* First of all, we'll populate the questions with the value of the minimum round acceptable, which in this case, will be 8 (all questions available).
	* Each question will take 2 rounds:
	* 	First: The main QUESTION
	*	Second: The complementary QUESTION
	*
	* So we'll have at least 16 questions.
	 */

	baseQuestions := 8

	questionsContainer := QuestionsContainer{}
	questions, err := makeShuffledQuestions(
		baseQuestions,
	)

	if err != nil {
		return make([]domain.Question, 0), err
	}
	questionsContainer.AddQuestion(questions...)

	// Minimum rounds = 15
	// If it's odd, we add 1 to the number of rounds to retrieve an integer rounds number.
	// So we'll have at least 16 rounds, but it will be interrupted at 15 it is the max rounds number
	if !utils.IntegerIsEven(int64(rounds)) {
		rounds++
	}

	if rounds > 16 {
		if err := makeExtraQuestions(rounds, baseQuestions, &questionsContainer); err != nil {
			return make([]domain.Question, 0), err
		}
	}

	return questionsContainer.Questions, nil
}

func makeShuffledQuestions(numQuestions int) ([]domain.Question, error) {
	return MakeQuestionsUsingWorkerPool(numQuestions)
}

func makeExtraQuestions(rounds uint, baseQuestions int, questionsContainer *QuestionsContainer) error {
	necessaryQuestion := int(rounds)/2 - baseQuestions

	additionalQuestions, err := makeShuffledQuestions(
		necessaryQuestion,
	)

	if err != nil {
		return err
	}

	/* Now populate the rest of the questions that the program will need, to achieve from 15 to 30 available rounds */
	questionsContainer.AddQuestion(additionalQuestions...)
	return nil
}
