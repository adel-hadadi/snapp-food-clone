package random

import (
	"math/rand"
)

func RandNum(length int) int {
	lowerBound := int64(1)
	for i := 1; i < length; i++ {
		lowerBound *= 10
	}
	upperBound := lowerBound * 10

	return int(rand.Int63n(upperBound-lowerBound) + lowerBound)
}
