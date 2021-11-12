package main

import "fmt"

func main()  {
	// new与make的区别
	//  1. new用于各种类型的内存分配，并返回指向这片内存的指针，make用于内建类型(map/slice/channel)的内存分配，make返回一个由初始值的T类型

	nameContainer := make(map[string]string)
	fmt.Println(nameContainer)

	nameContainer1 := new(map[string]string)
	fmt.Printf("%T \n", nameContainer1)
	fmt.Println(*nameContainer1)

	intVal := new(int)
	fmt.Printf("%T \n", intVal)
	fmt.Println(*intVal)

	strAddr := new(string)
	fmt.Printf("%T \n", strAddr)
	fmt.Println(*strAddr)
}

func newDemo()  {
	
}