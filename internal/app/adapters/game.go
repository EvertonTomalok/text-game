package adapters

import (
	"fmt"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/pkg/helpers"
	"github.com/evertotomalok/text-game/pkg/utils"
)

type UserInputAdapter struct {
}

func (a *UserInputAdapter) GetValidInput() uint {
	for {
		var option uint
		_, err := fmt.Scanln(&option)

		if err != nil {
			utils.ClearTerminal()
			fmt.Println("Please, type only 1 or 2...")
			continue
		}

		if option != 1 && option != 2 {
			utils.ClearTerminal()
			fmt.Println("Please, type only 1 or 2...")
			continue
		}

		return option
	}
}

type FlowController struct {
}

func (a *FlowController) FlowValidator(flow *config.Flow) error {
	return helpers.FlowValidator(flow)
}
