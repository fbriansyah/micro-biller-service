package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	seed := rand.NewSource(time.Now().UnixNano())
	rand.New(seed)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(1000, 100000)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomBillNumber() string {
	return fmt.Sprintf(
		"63%d%d%d%d",
		RandomInt(1000, 9999),
		RandomInt(1000, 9999),
		RandomInt(1000, 9999),
		RandomInt(1000, 9999),
	)
}
