package utils

import (
	"math/rand"
)

func getRandomNumber(min int, max int) int {
	return (rand.Int()*(max-min) - min)
}
