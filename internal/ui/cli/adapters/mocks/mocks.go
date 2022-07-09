package mocks

import (
	"github.com/evertotomalok/text-game/internal/app/config"
	"github.com/evertotomalok/text-game/internal/core/errs"
)

type FlowValidatorMockError struct {
}

func (a *FlowValidatorMockError) FlowValidator(flow *config.Flow) error {
	return &errs.InvalidRoundsNumber{}
}
