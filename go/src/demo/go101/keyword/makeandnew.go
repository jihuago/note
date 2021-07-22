package keyword

import "fmt"

/*
	* make的作用是初始化内置的数据结构，也就是切片、哈希表、channel
		make关键字的作用是创建切片、哈希表和channel等内置的数据结构
	* new的作用是根据传入的类型分配一片内存空间并返回指向这片内存空间的指针
 */
func KeywordDemo()  {
	//makeDemo()

	d := &Dog{Animal: Animal{"na"}}
	fmt.Println(d.Say())

	doubleArr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(doubleArr)

	var array [10]int
	slice := array[2:4]
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
}

func makeDemo()  {
	sli := make([]int, 1, 10)
	has := make(map[int]bool, 10)
	fmt.Printf("%T, %T\n", sli, has)
}

type Animal struct {
	Name string
}

type Dog struct {
	Animal
	color string
}

func (animal *Animal) Say() string {
	return "名字：" + animal.Name
}


