package main

import "fmt"

/*
语言切片
*/
func main() {

	var arr []int // 声明一个数组

	arr = []int{1, 2, 3}

	slice1 := make([]int, 3) // 用make来初始化一个长度为3的数组

	slice2 := []int{1, 2, 3}

	slice3 := slice2[:] // 初始化切片slice3,是数组slice2的引用

	s := slice3[1:3] // 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

	s := arr[3:] // 缺省endIndex时将表示一直到arr的最后一个元素

	s := arr[:5] // 缺省startIndex时将表示从arr的第一个元素开始

	s1 := s[1:3] // 通过切片s初始化切片s1

	s := make([]int, 3, 3) // 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

}

/**

 */
func append_copy() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
