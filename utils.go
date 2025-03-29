package utils

import (
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWZYZ1234567890-!?.,")

func GeneratePassword(length int) string {
	arr := make([]rune, length)
	for i := range arr {
		arr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(arr)
}
