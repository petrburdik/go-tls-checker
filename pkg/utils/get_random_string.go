package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(n int, runes string) string {
	var letterRunes = []rune(runes)
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
