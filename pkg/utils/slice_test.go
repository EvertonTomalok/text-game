package utils_test

import (
	"math/rand"
	"testing"

	. "github.com/evertotomalok/text-game/pkg/utils"
)

func TestShuffle(t *testing.T) {
	rand.Seed(0)

	baseItems := []uint{1, 2, 3, 4, 5, 6}
	items := make([]uint, len(baseItems))

	copy(items, baseItems)

	items = Shuffle(items)

	if len(baseItems) != len(items) {
		t.Fail()
	}

	for i, v := range items {
		if baseItems[i] != v {
			return
		}
	}

	t.Fail()
}
