package main

import (
	"data_element/app/controllers/learnGoWithTest"
	"testing"
)

// 为Hello函数编写测试
func TestHello(t *testing.T)  {
	name := "jack"
	got := learnGoWithTest.Hello(name)

	want := "Hello, jack"

	if got  != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
