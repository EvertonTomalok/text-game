package ports

import "github.com/evertotomalok/text-game/internal/application/config"

// go:generate mockery --name=InputGateway
type InputGateway interface {
	GetValidInput() uint
}

// go:generate mockery --name=FlowGateway
type FlowGateway interface {
	FlowValidator(flow *config.Flow) error
}
