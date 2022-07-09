package helpers

import "github.com/evertotomalok/text-game/internal/app/config"

func IncrementOneToComplexity(flow *config.Flow, incrementNum int) {
	flow.Mu.Lock()
	flow.DebugComplexity += incrementNum
	flow.Mu.Unlock()
}

func DecrementOneToComplexity(flow *config.Flow, decrementNum int) {
	flow.Mu.Lock()
	flow.DebugComplexity -= decrementNum
	flow.Mu.Unlock()
}
