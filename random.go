package gotools

import (
	"bytes"
	"math/rand"
	"time"
)

func RandomString(num int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < num; {
		if string(RandomInt(65, 90)) != temp {
			temp = string(RandomInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}
func RandomIntString(length int) string {
	var result bytes.Buffer
	var temp string
	var char string
	for i := 0; i < length; {
		char = string(RandomInt(48, 57))
		if char != temp {
			temp = char
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
