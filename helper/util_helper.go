package helper

import (
	"math/rand"
	"time"
)

func ContainOnList(list []string, value any) bool {
	for _, data := range list {
		if data == value {
			return true
		}
	}
	return false
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
