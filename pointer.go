package main

import "fmt"

/*
指针
*/
func main() {

	val := 3
	fmt.Println(&val) //返回变量内存地址

	var pointer *int /* 指向整型*/
	var fp *float32  /* 指向浮点型 */
	fmt.Println(pointer)
	fmt.Println(fp)

	/*
	   指针使用流程：

	   定义指针变量。
	   为指针变量赋值。
	   访问指针变量中指向地址的值。

	   在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

	*/
	var a int = 32 /*声明实际变量*/
	var ip *int    /*声明指针变量*/

	ip = &a /*指针变量的存储地址*/
	fmt.Printf("a变量的地址：%x\n", &a)

	/*指针变量的存储地址*/
	fmt.Printf("指针变量的存储地址：%x\n", ip)

	/*使用指针访问值*/
	fmt.Printf("ip 变量的值:%x\n", *ip)

	var ptr *int
	fmt.Printf("ptr的值%x\n", ptr)

	/*指向指针的指针*/
	var b int
	var p1 *int
	var p2 **int

	b = 3000
	p1 = &a
	p2 = &p1

	fmt.Printf("变量b的值：%d\n", b)

	fmt.Printf("p1指针：%d\n", p1)
	fmt.Printf("p2的指针:%d\n", p2)

	/*
		通过指针交换值
	*/
	z := 100
	x := 300

	fmt.Printf("交换前，z:%d,x:%d\n", z, x)
	use_swap(&z, &x)

	fmt.Printf("交换后，z:%d,x:%d", z, x)

}

func use_swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp

}

func use_swap2(a *int, b *int) {

	a, b = b, a

}

func use_swap3(a, b int) {

	a, b = b, a

}
