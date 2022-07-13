package cli

import (
	"fmt"

	"github.com/evertotomalok/text-game/internal/application/config"
	"github.com/evertotomalok/text-game/internal/application/shared"
	"github.com/evertotomalok/text-game/internal/core/domain"
	"github.com/evertotomalok/text-game/internal/core/errs"
	"github.com/evertotomalok/text-game/internal/core/ports"
	"github.com/evertotomalok/text-game/internal/ui/cli/builder"
	"github.com/evertotomalok/text-game/pkg/helpers"
	"github.com/evertotomalok/text-game/pkg/utils"
	"github.com/evertotomalok/text-game/pkg/utils/colors"
	log "github.com/sirupsen/logrus"
)

var GetInputUserAdapter ports.InputGateway
var FlowControllerAdapter ports.FlowGateway

func StartGame(flow *config.Flow) error {
	// Refactor and implements a CoR  design pattern where each node will be a step
	// https://refactoring.guru/design-patterns/chain-of-responsibility

	if err := FlowControllerAdapter.FlowValidator(flow); err != nil {
		log.Errorf(
			"%sIt was not possible to determine the number of rounds.%s",
			colors.RED,
			colors.END,
		)
		return &errs.InvalidRoundsNumber{}
	}

	questions, err := builder.CreateQuestions(flow.Rounds)
	if err != nil {
		log.Errorf("Something went wrong: Err: %+v", err)
		return err
	}

	if err = questionsHandler(flow, questions); err != nil {
		return err
	}

	getFlowResult(flow)

	return nil
}

func questionsHandler(flow *config.Flow, questions []domain.Question) error {
	for i, question := range questions {
		complemetaryQuestion, err := ProcessQuestion(question, flow)
		if err != nil {
			return err
		}

		if err = ProcessComplementaryQuestion(complemetaryQuestion, flow); err != nil {
			return err
		}

		if i+1 >= int(flow.Rounds) {
			break
		}

	}
	return nil
}

func ProcessQuestion(question domain.Question, flow *config.Flow) (domain.ComplementaryQuestion, error) {
	questionChoice := displayQuestion(question)
	var complemetaryQuestion domain.ComplementaryQuestion

	switch questionChoice {
	case 1:
		ComputeDebugComplexity(flow, question.FirstAlternative.DebugComplexityCalculate)
		complemetaryQuestion = shared.ComplementaryQuestions[uint(question.FirstAlternative.RedirectTo)]
	case 2:
		ComputeDebugComplexity(flow, question.SecondAlternative.DebugComplexityCalculate)
		complemetaryQuestion = shared.ComplementaryQuestions[uint(question.SecondAlternative.RedirectTo)]
	default:
		return complemetaryQuestion, &errs.InvalidQuestion{}
	}

	return complemetaryQuestion, nil
}

func ProcessComplementaryQuestion(complemetaryQuestion domain.ComplementaryQuestion, flow *config.Flow) error {
	complementaryChoice := displayComplementaryQuestion(complemetaryQuestion)

	fmt.Println()

	switch complementaryChoice {
	case 1:
		ComputeDebugComplexity(flow, complemetaryQuestion.FirstAlternative.DebugComplexityCalculate)
		fmt.Printf("\n-> %s\n\n", complemetaryQuestion.FirstAlternative.TransitionMessage)
	case 2:
		ComputeDebugComplexity(flow, complemetaryQuestion.SecondAlternative.DebugComplexityCalculate)
		fmt.Printf("\n-> %s\n\n", complemetaryQuestion.SecondAlternative.TransitionMessage)
	default:
		return &errs.InvalidComplementaryQuestion{}
	}

	fmt.Println()

	return nil
}

func displayQuestion(q domain.Question) uint {
	separator := utils.Separator("-", 150)

	println(separator)
	fmt.Printf("%sQuestion: %s%s\n", colors.YELLOW, colors.END, q.Text)
	println(separator)

	fmt.Printf("%s1 - %s %s\n", colors.BLUE, q.FirstAlternative.Text, colors.END)
	fmt.Printf("%s2 - %s %s\n\n", colors.GREEN, q.SecondAlternative.Text, colors.END)
	fmt.Printf("Choose one [1 - 2]: ")

	choice := GetInputUserAdapter.GetValidInput()
	utils.ClearTerminal()

	return choice
}

func displayComplementaryQuestion(c domain.ComplementaryQuestion) uint {
	separator := utils.Separator("-", 150)

	fmt.Println(separator)
	fmt.Printf("%sResult:%s %s \n", colors.YELLOW, colors.END, c.Text)
	fmt.Println(separator)

	fmt.Printf("%s1 - %s%s\n", colors.BLUE, c.FirstAlternative.Text, colors.END)
	fmt.Printf("%s2 - %s%s\n\n", colors.GREEN, c.SecondAlternative.Text, colors.END)
	fmt.Printf("Choose one [1 - 2]: ")

	choice := GetInputUserAdapter.GetValidInput()
	utils.ClearTerminal()

	return choice
}

func ComputeDebugComplexity(flow *config.Flow, complexity int) {
	if complexity >= 0 {
		helpers.IncrementOneToComplexity(flow, complexity)
	} else {
		abs := complexity * -1 // Get Absolute Value
		helpers.DecrementOneToComplexity(flow, abs)
	}
}

func getFlowResult(flow *config.Flow) {
	separator := utils.Separator("-", 80)

	fmt.Println(separator)

	if flow.DebugComplexity > 0 {
		fmt.Printf(
			"%sYour turn is over. You're fired because you couldn't handle the problems.%s\n",
			colors.RED,
			colors.END,
		)
	} else {
		fmt.Printf(
			"%sYou're a superhero! Thank you for saving the day!!!!!!%s\n",
			colors.GREEN,
			colors.END,
		)
	}

	fmt.Println(separator)

	log.Infof("[Complexity: %d]", flow.DebugComplexity)
}
