package main

import (
	"github.com/rayfarandi/fwg17-go-introduction/task1"
	"github.com/rayfarandi/fwg17-go-introduction/task2"
	"github.com/rayfarandi/fwg17-go-introduction/task3"
)

func main() {
	// task1

	task1.PrintSegitiga(2)

	// task 2

	customPassword := "abcd"
	strong := true
	task2.GenPass(customPassword, strong)

	// task3

	task3.LamaMenonton(7)

}
