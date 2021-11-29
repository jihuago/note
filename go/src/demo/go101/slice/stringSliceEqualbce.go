package main

import (
	"fmt"
)

func main() {
	a := []string{"b",}
	b := []string{"a",}
	//var a, b []string
	res := StringSliceEqualBCE(a, b)
	fmt.Println(res)
}

func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)] // 新版本go(1.15.6)已经优化，去掉 b = b[:len(a)]也不会边界检查 这种写法是为了优化边界检查从而提升运行时效率的。如果没有这一句，在运行时，Go语言每次都会对b[i]做边界检查
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
