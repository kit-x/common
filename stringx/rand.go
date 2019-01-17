package stringx

import (
	"math/rand"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString returns randomly generated string
func RandomString(l uint) string {
	rand.Seed(time.Now().UnixNano())
	s := make([]byte, l)
	for i := 0; i < int(l); i++ {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}
