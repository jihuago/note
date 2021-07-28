package _struct

import (
	"fmt"
	"strconv"
)

// interface是一组method签名的组合，通过interface来定义对象的一组行为

type Element interface {}
type List []Element

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func RunInterface()  {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"deni", 30}

	// Comma-ok断言， value, ok := ele.(T)
	for index, ele := range list {
		if value, ok := ele.(int); ok {
			fmt.Printf("list[%d] is an int and its values is %d\n", index, value)
			continue
		}

		// element.(type)语法不能在switch外的任何逻辑里面使用
		switch value := ele.(type) {
		case string:
			fmt.Printf("list[%d] is an string and its values is %s\n", index, value)
		}
	}
}