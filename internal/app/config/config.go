package config

import "sync"

type Flow struct {
	Mu              sync.Mutex
	Rounds          uint
	DebugComplexity int
}
