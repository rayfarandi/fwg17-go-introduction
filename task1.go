package main

import "fmt"

func PrintSegitiga(baris int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Task 1:", r)
		}
	}()

	if baris <= 0 {
		panic("baris tidak boleh negatif atau nol")
	}

	for i := baris; i > 0; i-- {
		for j := 0; j < baris-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
