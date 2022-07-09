package cli

import (
	"fmt"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/internal/core/ports"
	"github.com/evertotomalok/text-game/internal/ui/cli/builder"
	"github.com/evertotomalok/text-game/pkg/helpers"
)

var GetInputUser ports.InputInterface

func StartGame(flow *config.Flow) {
	if err := helpers.FlowValidator(flow); err != nil {
		fmt.Println("It was not possible to determine the number of rounds.")
		return
	}
	questions, _ := builder.CreateQuestions(flow.Rounds)
	for _, question := range questions {
		fmt.Printf("%+v\n\n", question)
	}
	fmt.Println(len(questions))

	_ = GetInputUser.GetValidInput()
}
