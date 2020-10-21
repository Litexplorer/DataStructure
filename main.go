package main

import (
	"./ArrayList"
	"fmt"
)

func main() {
	list := ArrayList.NewArrayList()
	list.Append("1")
	list.Append("2")
	list.Append("3")
	list.Append("4")
	list.Append("5")
	list.Append("6")
	fmt.Println(list.String())
	fmt.Printf("当前数组的长度是：【%d】\n", list.Size())

	//
	for i := 0; i < list.Size(); i++ {
		element, _ := list.Get(i)
		fmt.Printf("数组的第【%d】个元素是【%v】\n", i, element)
	}
}
