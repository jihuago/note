package controllers

import "fmt"

func DemoArr()  {
	printA()

	testArrValue()
}

func printA()  {
	a := [...]string{"a", "b", "c", "d"}

	for i, v := range a{
		fmt.Println("Array item", i, "is ", v)
	}
}

// Go语言中的数组是一种值类型，可以通过new()来创建
func testArrValue()  {
	var arr1 = new([5]int) // arr1 类型是*[5]int
	arr1[0] = 100

	var arr2 [5]int // 类型是[5]int

	arr2 = *arr1

	fmt.Printf("%T %T \n", arr1, arr2)
}