package main

import (
	"fmt"
)

func run_go(x int, y int) (int, int) {
	return y, x

}

func main() {

	expr := 99
	switch expr {
	case 99:
		println("99")
	case 20:
		println(88)

	default:
		println("default")

	}

	//var i, _ interface{}
	var i, _ interface{} = run_go(3, 5)

	switch i.(type) {

	case int:
		fmt.Print("this is int")
	case string:
		fmt.Print("this is string")
	case nil:
		fmt.Print("this is nil")
	default:
		fmt.Print("this is default")
	}

	//	fallthrough
	//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	z := 20
	switch z {
	case 20:
		fmt.Printf("this is %d\r\n", z)
		fallthrough
	case 50:
		fmt.Print("fallthough string ")
	case 70:
		fmt.Print("fallthough 70")
	default:
		fmt.Print("this is default")
	}

}
