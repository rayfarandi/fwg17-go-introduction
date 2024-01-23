package task2

import (
	"fmt"
	"math/rand"
)

func GeneratePassword(length int, low bool, med bool, strong bool, customPassword string) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var password []byte
	var charSource string

	if low {
		charSource += ""
	}
	if med {
		charSource += "0123456789"
	}
	if strong {
		charSource += "!@#$%^&*()_+=-"
	}
	charSource += charset

	if customPassword != "" {
		charSource += customPassword
	}

	for i := 0; i < length; i++ {
		randNum := rand.Intn(len(charSource))
		password = append(password, charSource[randNum])
	}
	return string(password)
}

func Task2(customPassword string, strong bool) {
	length := 12
	low := false
	med := false

	password := GeneratePassword(length, low, med, strong, customPassword)
	fmt.Println("Generated Password:", password)
}
