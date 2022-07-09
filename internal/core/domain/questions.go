package domain

import "github.com/evertotomalok/text-game/internal/app/config"

type ComplexityCalculator func(flow *config.Flow)

type QuestionAlternative struct {
	Text                     string
	RedirectTo               int
	DebugComplexityCalculate int
}

type ComplementaryQuestionAlternative struct {
	Text                     string
	DebugComplexityCalculate int
	TransitionMessage        string
}

type Question struct {
	ID                int
	Text              string
	FirstAlternative  QuestionAlternative
	SecondAlternative QuestionAlternative
}

type ComplementaryQuestion struct {
	ID                int
	Text              string
	FirstAlternative  ComplementaryQuestionAlternative
	SecondAlternative ComplementaryQuestionAlternative
}
