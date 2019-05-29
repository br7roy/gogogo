package main

import (
	"fmt"
	"time"
)

/*
并发
*/
func main() {

	for i, n := range []int{1, 2, 3} {
		go doWord(i, n)
	}

	doWord(9, 9)

	fmt.Println("main thread go")
	go say("hello")
	say("world")

}

func doWord(i int, n int) {
	fmt.Printf("doWork.index : %d,number : %d\n", i, n)
}

func say(s string) {
	for i := 10; i > 0; i-- {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}
