package cmd

import (
	"flag"
	"math/rand"
	"time"

	"github.com/evertotomalok/text-game/internal/application/adapters"
	"github.com/evertotomalok/text-game/internal/application/config"
	"github.com/evertotomalok/text-game/internal/ui/cli"
)

func Execute() {
	rand.Seed(time.Now().Unix())

	rounds := flag.Uint("rounds", 0, "Set number of rounds")
	flag.Parse()

	flow := config.Flow{
		Rounds:          *rounds,
		DebugComplexity: 0,
	}

	cli.GetInputUserAdapter = new(adapters.UserInputAdapter)
	cli.FlowControllerAdapter = new(adapters.FlowController)

	_ = cli.StartGame(&flow)
}
