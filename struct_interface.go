package main

import "fmt"

func main() {

	n := Nokia{}
	n.call()

	h := HuaWei{}
	h.call()

}

type MobileVendor interface {
	call()
}

type HuaWei struct {
}
type Nokia struct {
}

func (n Nokia) call() {
	fmt.Println("nokia shouji call")
}

func (h HuaWei) call() {
	fmt.Println("huawei shouji call")
}
