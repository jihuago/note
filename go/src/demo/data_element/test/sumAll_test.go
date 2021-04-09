package main

import (
	"data_element/app/controllers/learnGoWithTest/integers"
	"fmt"
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T)  {
	got := integers.SumAll([]int{1, 2}, []int{0, 9})

	want := []int{3, 9}
	//want := "bob"

	//if got != want { // Go中不能对切片使用等号运算符。解决方法：1. 写一个而寒暑迭代每个元素来检查他们的值
	// 2. 使用reflect.DeepEqual，判断两个变量是否相等
	// 注意：reflect.DeepEqual不是类型安全的，所以有时候会发生比较怪异的行为。
	if ! reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll()  {
	got := integers.SumAll([]int{1, 2}, []int{3, 7})

	fmt.Println(got)

	// Output: [3 10]
}