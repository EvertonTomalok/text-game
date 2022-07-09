package cli

import (
	"fmt"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/internal/core/errs"
	"github.com/evertotomalok/text-game/internal/core/ports"
	"github.com/evertotomalok/text-game/internal/ui/cli/builder"
	log "github.com/sirupsen/logrus"
)

var GetInputUserAdapter ports.InputInterface
var FlowControllerAdapter ports.FlowInterface

func StartGame(flow *config.Flow) error {
	if err := FlowControllerAdapter.FlowValidator(flow); err != nil {
		log.Error("It was not possible to determine the number of rounds.")
		return &errs.InvalidRoundsNumber{}
	}
	questions, err := builder.CreateQuestions(flow.Rounds)
	if err != nil {
		log.Errorf("Something went wrong: Err: %+v", err)
		return err
	}

	for _, question := range questions {
		fmt.Printf("%+v\n\n", question)
	}
	fmt.Println(len(questions))

	_ = GetInputUserAdapter.GetValidInput()
	return nil
}
