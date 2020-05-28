package main

import "fmt"

/*
结构体
类型
*/

type Book struct {
	title  string
	bookId string
	author string
}

func main() {

	fmt.Println(Book{"北方的山", "007", "鲁班"})
	/*使用 key=>value形式*/
	fmt.Println(Book{title: "声东击西", bookId: "008", author: "孙膑"})

	/*访问结构体成员*/
	var book1 Book
	//var book2 Book
	book1.author = "克鲁斯托夫"

	/*结构体指针*/
	var bookPtr *Book
	bookPtr = &book1

	/*使用结构体指针访问结构体成员，使用 "." 操作符：*/
	fmt.Println(bookPtr.author)

	rect := Rect{float64(200), float64(300), float64(20), float64(30)}

	r := rect.Area()
	fmt.Println(r)

}

type Rect struct {
	//定义矩形类
	x, y          float64 //类型只包含属性，并没有方法
	width, height float64
}

func (r *Rect) Area() float64 { //为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
	return r.width * r.height //方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
}
