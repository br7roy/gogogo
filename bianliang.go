package main

import (
	"fmt"
)

// 以下几种类型变量零值为nil
var (
	//a *int
	//a []int
	//a map[string]int
	//a chan int
	a func(string) int
	//a error // error 是接口
)

func main() {
	println(a)
	if a == nil {
		fmt.Print("a is nil")
		fmt.Print(a)
	}
}
