package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomFloat generates a random float between min and max
func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomString generates a random string of length n
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return randomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() float64 {
	return randomFloat(0, 500)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) *int64 {
	ran := min + rand.Int63n(max-min+1)
	return &ran
}
