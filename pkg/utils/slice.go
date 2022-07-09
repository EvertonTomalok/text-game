package utils

import "math/rand"

func Shuffle(item []uint) []uint {
	for i := len(item) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		item[i], item[j] = item[j], item[i]
	}
	return item
}
