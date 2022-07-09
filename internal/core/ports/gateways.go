package ports

import "github.com/evertotomalok/text-game/internal/app/config"

// go:generate mockery --name=InputInterface
type InputInterface interface {
	GetValidInput() uint
}

// go:generate mockery --name=FlowInterface
type FlowInterface interface {
	FlowValidator(flow *config.Flow) error
}
