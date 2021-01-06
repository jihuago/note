package main

import (
	"fmt"
	"os"
	"data_element/common"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
}

func main()  {
	// 数据类型

/*	// 整型
	var number1 int16
	number1 = 12

	var number2 int8
	number2 = 12

	// 不同类型的值不能放一起比较
	//res1 := (number2 == number1)

	// 不同类型的整型值不能直接进行算数运算 int8()转换类型为int8
	res := number2 + int8(number1)
	res ++

	// 各种类型的整型变量都可以直接与字面常量进行比较
	res4 := res == 2

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(res)
	fmt.Println(res4)*/

	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	firstChr, length := common.GetStringFirstChar("啊test")
	fmt.Println(firstChr)
	fmt.Println(length)

	str := "啊test"
	fmt.Println(len(str))

	common.TestDataTranfer()

	common.ArrMethod()

}
