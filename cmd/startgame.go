package cmd

import (
	"flag"

	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/internal/ui/cli"
	"github.com/evertotomalok/text-game/internal/ui/cli/adapters"
)

func Execute() {
	rounds := flag.Uint("rounds", 0, "Set number of rounds")
	flag.Parse()

	flow := config.Flow{
		Rounds:          *rounds,
		DebugComplexity: 0,
	}

	inputUserAdapter := new(adapters.UserInputAdapter)
	cli.GetInputUser = inputUserAdapter

	cli.StartGame(&flow)
}
