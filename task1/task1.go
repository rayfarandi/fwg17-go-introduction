package task1

import "fmt"

func PrintSegitiga(baris int) {
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
