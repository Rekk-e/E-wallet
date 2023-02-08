package utils

import (
	"math/rand"
	"strings"
	"time"
)

// Генерация случайного адреса для кошельков
func СreateAddress() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"abcdefghijklmnopqrsstuvwxyz" +
			"0123456789")
	length := 64
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
