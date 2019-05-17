package main

import (
	"fmt"
)

var (
	Name string
	Age  uint
)

func main() {
	print("hello world!")
	var b, c int = 1, 2
	fmt.Print(b, c)
	var intVal int
	print(intVal)
	f := "go"
	print(f)
	_, r2, r3 := numbers()

	print(r2, r3)
}

func numbers() (int, int, string) {
	a, b, c := 1, 2, "hehe"
	return a, b, c
}
