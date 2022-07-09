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

type FlowValidatorMock struct {
}

func (a *FlowValidatorMock) FlowValidator(flow *config.Flow) error {
	return nil
}

type InputMock struct {
}

func (a *InputMock) GetValidInput() uint {
	return uint(1)
}
