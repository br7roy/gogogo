package main

import "fmt"

/*
	channel (通道)
*/
func main() {

	//通道（channel）是用来传递数据的一个数据结构。
	//声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：

	ch := make(chan int)

	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go addAll(i[:len(i)/2], ch)
	go addAll(i[len(i)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

	//	Go 遍历通道与关闭通道
	//Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：
	/*v, ok := <-ch*/

	//如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭。

	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)

	}

}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func addAll(n []int, c chan int) {
	sum := 0
	for _, n := range n {
		sum += n
	}
	c <- sum
}
