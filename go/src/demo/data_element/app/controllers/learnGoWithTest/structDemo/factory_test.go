package structDemo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewFile(t *testing.T) {
	t.Run("test new object", func(t *testing.T) {
		f := NewFile(10, "./test.txt")

		// 获取结构体的实例占用了多少内存
		size := unsafe.Sizeof(f)

		fmt.Println(size, f)
	})

	t.Run("test force factory method", func(t *testing.T) {
		m := NewMatrix(2)

		fmt.Println(m)
	})
}
