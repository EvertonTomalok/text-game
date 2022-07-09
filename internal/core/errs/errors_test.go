package errs_test

import (
	"errors"
	"testing"

	. "github.com/evertotomalok/text-game/internal/core/errs"
)

func TestNotNumberError(t *testing.T) {
	err := new(NotNumberError)

	target := &NotNumberError{}
	if !errors.As(err, &target) {
		t.Fail()
	}

}

func TestNotInvalidRoundsNumber(t *testing.T) {
	err := new(InvalidRoundsNumber)

	target := &InvalidRoundsNumber{}
	if !errors.As(err, &target) {
		t.Fail()
	}

}
