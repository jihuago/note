package errHanlde

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	自定义包中的错误处理和panicking
	* 这是所有自定义包实现者应该遵守的最佳实践
		1. 在包内部，总是应该从panic中recover：不允许显式的超出包范围的panic()
		2. 向包的调用者返回错误值，而不是panic
	* 在包内部，特别是在非导出函数中有很深层次的嵌套调用时，将panic转换成error来告诉调用方为何出错，这提高了代码可读性
	*
*/

// parse包用来将输入的字符串解析为整数切片
type ParseError struct {
	Index int
	Word string
	Err error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse:error parsing %q as int", e.Word)
}

func Parse(input string) (numbers []int, err error)  {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok != true {
				err = fmt.Errorf("pkg:%v", err)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2number(fields)
	
	return
}

func fields2number(field []string) (numbers []int)  {
	if len(field) == 0 {
		panic("no words to parse")
	}

	for idx, field := range field {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}

		numbers = append(numbers, num)
	}

	return
}