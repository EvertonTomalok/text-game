package helpers

import (
	"fmt"

	"github.com/evertotomalok/text-game/internal/app/config"
	e "github.com/evertotomalok/text-game/internal/core/errs"
	"github.com/evertotomalok/text-game/pkg/utils"
	"github.com/evertotomalok/text-game/pkg/utils/colors"
)

func FlowValidator(flow *config.Flow) error {
	utils.ClearTerminal()
	if flow.Rounds < 15 || flow.Rounds > 30 {
		if err := loopToCollectRoundsNumber(flow); err != nil {
			return err
		}
	}

	printRoundsInfo(flow)

	return nil
}

func loopToCollectRoundsNumber(flow *config.Flow) error {
	// TODO get io.Reader as param, to be easier to test
	for _, try := range []int{1, 2, 3} {
		fmt.Printf("%d - The number of rounds must be a NUMBER between 15 and 30. Chose a value: ", try)
		var rounds uint
		_, err := fmt.Scanln(&rounds)

		if err != nil {
			fmt.Println("Please, type only NUMBERS...")
			continue
		}

		if rounds >= 15 && rounds <= 30 {
			flow.Rounds = rounds
			return nil
		}
	}
	return new(e.NotNumberError)
}

func printRoundsInfo(flow *config.Flow) {
	utils.ClearTerminal()

	separator := utils.Separator("-", 36)

	fmt.Println(separator)
	fmt.Printf(
		"| %sNice. We'll play with %d rounds.%s |\n",
		colors.GREEN,
		flow.Rounds,
		colors.END,
	)
	fmt.Printf("%s \n\n\n", separator)
}
