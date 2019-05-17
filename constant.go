package main

import "fmt"

// 常量

const (
	a = iota
	b
	c
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {

	println(a, b, c)

	print(i, j, k, l)

	for true {

		fmt.Print("这是一个死循环")

	}

}
