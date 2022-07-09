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

	var questions []domain.Question = suffleQuestions()

	if !utils.IntegerIsEven(int64(rounds)) {
		rounds++
	}

	baseQuestions := 8
	necessaryQuestion := rounds/2 - uint(baseQuestions)
	var additionalQuestions []domain.Question = suffleQuestions()

	for i := 0; i < int(necessaryQuestion); i++ {
		questions = append(questions, additionalQuestions[i])
	}

	return questions, nil
}

func suffleQuestionsNumbers() []uint {
	numbers := make([]uint, len(shared.MainQuestions))

	for i := 1; i < len(shared.MainQuestions); i++ {
		numbers[i] = uint(i)
	}

	return utils.Shuffle(numbers)
}

func suffleQuestions() []domain.Question {
	randomNumbers := suffleQuestionsNumbers()

	var questions []domain.Question

	for _, v := range randomNumbers {
		questions = append(questions, shared.MainQuestions[v])
	}

	return questions
}
