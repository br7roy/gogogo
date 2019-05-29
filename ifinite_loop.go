package main

import "fmt"

/*
递归
*/
func main() {

	fmt.Println(loop(3))

}
func loop(n int64) int64 {
	if n > 0 {
		sum := n * loop(n-1)
		return sum
	}
	return 1
}
