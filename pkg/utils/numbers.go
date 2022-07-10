package utils

func IntegerIsEven(number int64) bool {
	return number%2 == 0
}

func URange(start uint, end uint) []uint {
	numbers := make([]uint, end)

	for i := start; i < end; i++ {
		numbers[i] = uint(i)
	}

	return numbers
}
