package main

import (
	"fmt"
	"math/rand"
)

func GeneratePassword(length int, low bool, med bool, strong bool, customPassword string) string {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Panic Task 2:", p)
		}
	}()

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var password []byte
	var charSource string

	if length <= 0 {
		panic("Panjang password harus lebih besar dari 0")
	}
	if low {
		charSource += ""
	}
	if med {
		charSource += "0123456789"
	}
	if strong {
		charSource += "0123456789!@#$%^&*()_+=-"
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

func GenPass(customPassword string, strong bool, length int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Panic Task 2:", p)
		}
	}()

	low := false
	med := false

	password := GeneratePassword(length, low, med, strong, customPassword)
	fmt.Printf("Custom Password yang Anda masukan adalah :  %s ", customPassword)
	fmt.Println()
	fmt.Println("Generated Password:", password)
}
