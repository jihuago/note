package common

import (
	"fmt"
	"reflect"
)

func GetStringFirstChar(str string) (string, int) {
	// 将string转为字符类型数组
	runStr := []rune(str)

	return string(runStr[0]), len(runStr)
}

func TestDataTranfer() {
	// 整型、浮点型转换
	v1 := 99.99
	v2 := int(v1)

	fmt.Println("浮点转整型使用int():", v2)

	// 将整型转换成浮点型
	intNumber := 99
	floatNumber := float64(intNumber)
	fmt.Println(reflect.TypeOf(floatNumber))
}


// 数组操作
func ArrMethod() {
	// 定义一个数组
	arr := [2]int{2, 3}

	// 打印
	fmt.Println(arr)

	// 遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println("下标：", i, "元素：", arr[i])
	}

	// range 遍历
	for i, v := range arr {
		fmt.Println(i, "=>", v)
	}

	// 数组切片
	months := [...]string{
		"January", "February", "March", "April"}

	q2 := months[1:3]
	all := months[:]

	fmt.Println(q2)
	fmt.Println(all)

}




