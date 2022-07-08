package domain

import (
	"github.com/evertotomalok/text-game/internal/app/config"
)

type Question struct {
	ID               int
	Text             string
	FirstAlternative struct {
		Text                     string
		RedirectTo               int
		DebugComplexityControler func(flow *config.Flow)
	}
	SecondAlternative struct {
		Text                     string
		RedirectTo               int
		DebugComplexityControler func(flow *config.Flow)
	}
}
