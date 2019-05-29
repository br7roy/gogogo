package main

import (
	"fmt"
)

/**
range函数
*/
func main() {
	nums := []int{1, 2, 10, 4, 5}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	kvs := map[string]string{"1": "asdaw", "2": ";lkjl"}

	//range也可以用在map的键值对上
	for k, v := range kvs {
		fmt.Printf("k:%s=v:%s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
