package main

func main() {

	// Task 1
	PrintSegitiga(5)
	// jika ingin memicu panic
	// PrintSegitiga(-2)

	// Task 2

	customPassword := "abcd"
	strong := true
	// length := 12
	// jika ingin memicu panic
	length := -2
	GenPass(customPassword, strong, length)

	// Task 3

	LamaMenonton(7)
	// jika ingin memicu panic
	// LamaMenonton(99)

}
