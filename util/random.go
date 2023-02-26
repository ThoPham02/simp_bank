package util

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	letters    = "abcdefghijklmnopqrstuvwxyz"
	currencies = []string{"USD", "EUR", "GBP"}
)

// This is a special function will be called automatically when the package is first created
// If don't have this function, random number is the same value for every run.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer between min and max
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// RandomString returns a random string has length n
func RandomString(n int64) string {
	len := len(letters)
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len)]
	}
	return string(b)
}

// RandomOwner returns a random an owner has length n
func RandomOwner(n int64) string {
	return RandomString(n)
}

func RandomBallance(n int64) int64 {
	return RandomInt(0, n)
}

func RandomCurrency() string {
	len := len(currencies)
	return currencies[rand.Intn(len)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(8))
}
