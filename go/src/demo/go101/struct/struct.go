package _struct

import (
	"encoding/json"
	"fmt"
)

// 每个结构体字段在声明可以被指定一个标签，实践中标签应该表示成用空格分隔的键值对刑事，并且用``表示，而键值对中的值使用解释型字面形式""表示
// 每个字段标签的目的取决于具体应用
// 把字段标签当成字段注释来使用不是一个好主意
// 标签是结构体的元信息，可以在运行时通过反射机制读取处理
type Persion struct {
	Name string `json:"姓名" myfmt:"s1"`
	sex int // 私有字段不能被json包访问
}

type student struct {
	name string
	age int
}

func StructDefine()  {
	p := &Persion{"jack", 1}
	//fmt.Println(p.Name)
	data, err := json.Marshal(p)
	if err != nil {

	}
	fmt.Printf("%s \n", data)

	demoStruct()
}

func demoStruct()  {
	m := make(map[string]*student)
	stus := []student{
		{name: "pp", age: 11},
		{name: "p2", age: 12},
	}

	// 因为&stu地址是固定的
	for _, stu := range stus {
		m[stu.name] = &stu
		fmt.Printf("%p \n", &stu)
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.age)
	}
}
