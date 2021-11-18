package main

import (
	"testing"
)

func BenchmarkStringPlus(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		StringPlus()
	}


}

func StringPlus() string {
	var s string
	s += "昵称" + ":"  + "\n"

	return s
}
