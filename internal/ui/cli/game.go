package cli

import (
	"fmt"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/pkg/helpers"
)

func StartGame(flow *config.Flow) {
	if err := helpers.FlowValidator(flow); err != nil {
		fmt.Println("It was not possible determine the round numbers.")
		return
	}

}
