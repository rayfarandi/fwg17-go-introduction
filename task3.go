package main

import "fmt"

func PilihFilem(duration int) string {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println("Panic Task 3:", f)
		}
	}()

	data := [...]int{1, 7, 3, 4, 8, 9}
	var result string
	for i, filmPertama := range data {
		for j, filmKedua := range data {
			if i != j && filmPertama+filmKedua == duration {
				result = fmt.Sprintf(`Rekomendasi Film dengan durasi penerbangan %d jam
Film ke-%d dengan durasi %d jam dan film ke-%d dengan durasi %d jam`, duration, i+1, filmPertama, j+1, filmKedua)
				return result
			}
		}
	}
	panic(fmt.Sprintf("Tidak ada rekomendasi film yang pas untuk penerbangan dengan durasi %d jam", duration))
}

func LamaMenonton(durasiPenerbangan int) {
	fmt.Printf("%s", PilihFilem(durasiPenerbangan))
}
