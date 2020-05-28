package main

import "fmt"

/*
数组
*/
func main() {

	sz := [...]int{1, 2, 3, 4, 5}
	sz2 := []int{1, 2, 3, 4, 5}
	fmt.Print(sz)
	fmt.Println(sz2)
	ints := append(sz2, 6)
	fmt.Println(ints)

}
