package main

import "fmt"

// 常量
//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式
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
