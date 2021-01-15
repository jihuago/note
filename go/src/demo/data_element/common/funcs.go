package common

import "fmt"

// go函数
/*

	函数可以以声明的方式被使用，作为一个函数类型。这里，不需要函数体{}
		type binOp func(int, int) int

	Go没有泛型(generic)的概念，也就是说Go不支持那种支持多种类型的函数。
		不过在大部分情况下可以通过接口，特别是空接口与类型选择来实现类似的功能。

	如果一个函数需要返回值，那么这个函数里面的每个代码分支都要有return 语句

	* 按值传递 or 按引用传递
		* 如果希望函数可以直接修改参数的值，而不是对参数的副本进行操作，需要将参数的地址（变量名前面添加&符合，比如&variable）传递给
函数，这就是按引用传递，比如Function(&arg1)，此时传递给函数的是一个指针。
		* 如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；可以通过这个指针的值来修改这个值所
指向的地址上的值（指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。）
		* 在函数调用时，像切片、字典、接口、通道这样的引用类型都是默认使用引用传递（即使没有显示指出指针）
		* 如果一个函数需要返回4，5个值，我们可以传递一个切片给函数或者传递一个结构体。因为传递一个指针允许直接修改变量的值，消耗也更少。
		* 使用非命名返回值是很糟的编程习惯，尽量使用命名返回值：会使代码更清晰、更简短、同时更加容易读懂
			return 2 * input // 这个习惯不好


*/

// 返回多个未命名值
/*func Test(name string, age int) (string, int) {
	return name, age
}*/

// 命名的返回值
func mult_returnval(number1, number2 int) (sum, product, diff int) {

	sum, product, diff = number1 + number2, number1 * number2, number1 - number2
	return
}

// 传递变长参数
/*func min(args ...int) {

}*/

/*
	将函数作为参数
		函数可以作为其他函数的参数进行传递，然后在其他函数内调用执行，一般称之为回调。
*/
func add(a, b int) (sum int)  {
	sum = a + b
	return
}

func callback(f func(i, j int) (int)) int {
	return f(1, 2)
}

func Run()  {
	res := callback(add)
	fmt.Println(res)
}


