package controllers

import "fmt"

func DemoInterview()  {
	// 结构体比较问题
	// 结构体比较规则1：只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关
	// 结构体比较规则2：结构体时相同的，但是结构体属性中有不可以比较的类型，如map
/*	sn1 := struct {
		age int
		name string
	}{
		age: 11,
		name: "jack",
	}

	sn2 := struct {
		age int
		name string
	}{
		age: 11,
		name: "jack",
	}

	if sn1 == sn2 {
		fmt.Println("sn1==sn2")
	}*/

	/*
		nil 可以用作interface,function,pointer,map,slice和channel的空值，但是如果不特别指定，Go不能识别类型。
	*/
	//demoReturnNil()

	// 在Golang中，常量归属于全局区，常量为存放数值字面值单位，即不可修改。

	// 切片的初始化与追加  切片追加，make初始化均为0
	appendSlice()

	testStruct()

	testMapRange()
}

func demoReturnNil() (string, bool)  {
	//return nil, false
	return "", false
}

func appendSlice() {
	s := make([]int, 10)
	s = append(s, 1, 2, 3)

	// 两个slice在append，需要进行将第二个slice进行...打散再拼接
	s2 := []int{4, 5}
	s2 = append(s, s2...)

	fmt.Println(s2)
}

// new和make的区别
/*
	二者都是内存的分配，但是make只用于slice、map以及channel的初始化；而new用于类型的内存分配，并且内存置为零
make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
*/

/*
	map[string]Student的value是一个Student结构值，所以当list["student"] = student是一个值拷贝过程。而list["student"]
则是一个值引用。那么值引用的特点是 只读。所以对list["student"].Name = "LDB"的修改是不允许的
*/
type Student struct {
	Name string
}
var list map[string]*Student

func testStruct()  {
	list = make(map[string]*Student)
	student := Student{"Aceld"}

	//list["student"] = student
	//list["student"].Name = "LDB" // 错误

	list["student"] = &student
	list["student"].Name = "LDB"

	fmt.Println(list["student"].Name)

}

type student struct {
	Name string
	Age int
}
/*
	foreach 中，stu是结构体的一个拷贝副本，所以m[stu.Name] = &stu 实际上一直指向同一个指针，最终该指针的值为遍历的最后一个struct的值拷贝
 */
func testMapRange()  {
	// 定义map
	m := make(map[string]*student)

	stus := []student{
		{Name: "zhou", Age: 20},
		{Name: "li", Age: 22},
		{Name: "wang", Age: 28},
	}

	// range陷阱
	// range与指针一起使用时，经常会碰到值都变成了最后一次循环的值
	// 原因：在GO的for ... range 循环中，Go始终使用值拷贝的方式代替被遍历的元素本身，就是
	// for .. range 中那个stu是一个值拷贝，而不是元素本身。&stu实际上只是取到了stu这个临时变量的地址，
	// 而非list中真正被遍历到的某个元素的地址。m[stu.Name] 被填充了三个相同的地址，都是&stu地址
	for _, stu := range stus {
		fmt.Printf("%T %v\n", &stu, &stu)
		m[stu.Name] = &stu
	}

	fmt.Println(m)
/*	for i := 0; i < len(stus); i++ {
		fmt.Printf("%T %v\n", &stus[i], &stus[i])
	}*/
}

