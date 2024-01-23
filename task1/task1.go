package task1

import "fmt"

func PrintSegitiga(tinggi int) {
	for i := tinggi; i > 0; i-- {
		for j := 0; j < tinggi-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
