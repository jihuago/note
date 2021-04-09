package integers

import (
	"fmt"
	"testing"
)

/*
	1. 进入到integers目录
	2. go test

 */

func TestAdder(t *testing.T)  {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// 通过添加这段代码，示例将出现在godoc的文档中，这将使你的代码更容易理解
// godoc -http=:6060
// 如果删除下面注释( Output: 6)，示例函数将不会执行。
func ExampleAdd()  {
	sum := Add(1, 5)
	fmt.Println(sum)

	// Output: 6
}

// 利用注释为函数添加文档，这些都将出现在Go Doc中，就像你查看标准库的文档一样
// Add takes two integers and returns the sum of them
func Add(x, y int) (sum int) {
	sum = x + y

	return
}
