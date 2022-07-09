package cli

import (
	"bytes"
	"fmt"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/internal/app/shared"
	"github.com/evertotomalok/text-game/internal/core/domain"
	"github.com/evertotomalok/text-game/internal/core/errs"
	"github.com/evertotomalok/text-game/internal/core/ports"
	"github.com/evertotomalok/text-game/internal/ui/cli/builder"
	"github.com/evertotomalok/text-game/pkg/helpers"
	"github.com/evertotomalok/text-game/pkg/utils"
	log "github.com/sirupsen/logrus"
)

var GetInputUserAdapter ports.InputInterface
var FlowControllerAdapter ports.FlowInterface

func StartGame(flow *config.Flow) error {
	// Refactor and implements a CoR  design pattern where each node will be a step
	// https://refactoring.guru/design-patterns/chain-of-responsibility

	if err := FlowControllerAdapter.FlowValidator(flow); err != nil {
		log.Error("It was not possible to determine the number of rounds.")
		return &errs.InvalidRoundsNumber{}
	}

	questions, err := builder.CreateQuestions(flow.Rounds)
	if err != nil {
		log.Errorf("Something went wrong: Err: %+v", err)
		return err
	}

	questionsHandler(flow, questions)
	getFlowResult(flow)

	return nil
}

func questionsHandler(flow *config.Flow, questions []domain.Question) {
	for _, question := range questions {
		questionChoice := displayQuestion(question)
		var complemetaryQuestion domain.ComplementaryQuestion

		switch questionChoice {
		case 1:
			computeDebugComplexity(flow, question.FirstAlternative.DebugComplexityCalculate)
			complemetaryQuestion = shared.ComplementaryQuestions[uint(question.FirstAlternative.RedirectTo)]
		case 2:
			computeDebugComplexity(flow, question.SecondAlternative.DebugComplexityCalculate)
			complemetaryQuestion = shared.ComplementaryQuestions[uint(question.SecondAlternative.RedirectTo)]
		}

		complementaryChoice := displayComplementaryQuestion(complemetaryQuestion)

		fmt.Println()

		switch complementaryChoice {
		case 1:
			computeDebugComplexity(flow, complemetaryQuestion.FirstAlternative.DebugComplexityCalculate)
			fmt.Println(complemetaryQuestion.FirstAlternative.TransitionMessage)
		case 2:
			computeDebugComplexity(flow, complemetaryQuestion.SecondAlternative.DebugComplexityCalculate)
			fmt.Println(complemetaryQuestion.SecondAlternative.TransitionMessage)
		}

		fmt.Println()
	}
}

func displayQuestion(q domain.Question) uint {
	fmt.Println("Question: ", q.Text)
	fmt.Println("1 - ", q.FirstAlternative.Text)
	fmt.Println("2 - ", q.SecondAlternative.Text)
	fmt.Printf("Chose one [1 - 2]: ")

	choice := GetInputUserAdapter.GetValidInput()
	utils.ClearTerminal()

	return choice
}

func displayComplementaryQuestion(c domain.ComplementaryQuestion) uint {
	fmt.Println()
	fmt.Println(c.Text)
	fmt.Println("1 - ", c.FirstAlternative.Text)
	fmt.Println("2 - ", c.SecondAlternative.Text)
	fmt.Printf("Chose one [1 - 2]: ")

	choice := GetInputUserAdapter.GetValidInput()
	utils.ClearTerminal()

	return choice
}

func computeDebugComplexity(flow *config.Flow, complexity int) {
	if complexity > 0 {
		helpers.IncrementOneToComplexity(flow, complexity)
	} else {
		helpers.DecrementOneToComplexity(flow, complexity*-1)
	}
}

func getFlowResult(flow *config.Flow) {
	var buffer bytes.Buffer
	for i := 0; i < 150; i++ {
		buffer.WriteString("-")
	}

	fmt.Println(buffer.String())
	if flow.DebugComplexity > 0 {
		fmt.Printf("Your turn is over. You're fired because you couldn't handle with the problems. [Complexity: %d]\n", flow.DebugComplexity)
	} else {
		fmt.Printf("You're a super hero! Thank you for saving the day!!!!!! [Complexity: %d]\n", flow.DebugComplexity)
	}
	fmt.Println(buffer.String())
}
