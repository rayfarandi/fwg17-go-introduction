package main

import (
	"github.com/rayfarandi/fwg17-go-introduction/task1"
	"github.com/rayfarandi/fwg17-go-introduction/task2"
)

func main() {
	// task1

	task1.PrintSegitiga(5)

	// task 2

	customPassword := "abcd"
	strong := true
	task2.Task2(customPassword, strong)

}
