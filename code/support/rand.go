package support

import (
	"math/rand"
)

func GenerateRandomInt() int {
	randomNumber := rand.Intn(5)

	return randomNumber
}