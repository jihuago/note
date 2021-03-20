package controllers

import (
	"fmt"
	"strings"
)

// strings
func DemoStrings()  {

	/*
		字符串是UTF-8字符的一个序列（当字符为ASCII码时占用1个字节，其他字符根据需要占用2-4个字节）
	 */
	s1 := `\naaa\t` // 非解析字符串，使用反引号括起来，``中的原样输出

	// 字符串拼接符 +
	// 在循环中使用加号+拼接字符串并不是最高效的做法，更好的方法是使用函数strings.Join()，更高效的是bytes.Buffer
	s1 += "this is a test"
	fmt.Println(s1)

	res := strings.Count("ba1", "")
	fmt.Println(res)

	// 查找
	// 字符串包含关系
	content := "this is a test 这是一个测试"
	bool := strings.Contains(content, "测试2")
	fmt.Println(bool)

	// 修剪字符串
	res1 := strings.TrimSpace(" abc ")
	fmt.Println(res1)

	// strings.Fields()利用1个或多个空白符号作为动态长度的分隔符将字符串分割成若干小块
	sli := strings.Fields("a b c  this is")
	for _, v := range sli {
		fmt.Println(v)
	}

	// strings.Join()将元素类型为string的clice使用分隔符号来拼接成一个字符串
	s2 := []string{
		"a",
		"b",
	}
	re := strings.Join(s2, "-")
	fmt.Println(re)


}
