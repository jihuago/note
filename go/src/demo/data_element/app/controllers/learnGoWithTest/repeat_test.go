package learnGoWithTest

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

// go test -v    可以示例代码测试
func ExampleRepeat()  {
	repeated := Repeat("a", 2)

	fmt.Println(repeated)

	// Output: aa
}

// 基准测试： 在Go中编写基准测试是Go语言的一个一级特性，它与编写测试非常相似。
// go test -bench=. 来运行基准测试
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 50)
	}
}
