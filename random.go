package gotools

import (
	"bytes"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomString(num int) string {
	var result bytes.Buffer
	for i := 0; i < num; i++ {
		result.WriteRune(rune(RandomInt(65, 90)))
	}
	return result.String()
}

func RandomIntString(length int) string {
	var result bytes.Buffer
	for i := 0; i < length; i++ {
		result.WriteRune(rune(RandomInt(48, 57)))
	}
	return result.String()
}

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
