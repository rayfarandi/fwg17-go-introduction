package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genPassword(password string, level string) string {

	var tambahanKarakter string
	switch level {
	case "low":
		tambahanKarakter = "0123456789"
	case "medium":
		tambahanKarakter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	case "strong":
		tambahanKarakter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

	default:
		return "Level Tidak ada"
	}

	var newPassword []rune
	newPassword = append(newPassword, []rune(password)...)

	for i := len(newPassword); i < 9; i++ {
		newPassword = append(newPassword, []rune(tambahanKarakter)[rand.Intn(len(tambahanKarakter))])
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newPassword), func(i, j int) {
		newPassword[i], newPassword[j] = newPassword[j], newPassword[i]
	})

	return string(newPassword)

}

func main() {
	password := "abcd"
	level := "strong"

	newPassword := genPassword(password, level)
	fmt.Println("password baru :", newPassword)
}
