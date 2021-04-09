package learnGoWithTest

import (
	"data_element/app/controllers/learnGoWithTest/integers"
	"testing"
)

/*
		numbers := [5]int{1, 2, 3, 4, 5}
		numbers := [...]int{1, 2, 3}

		1. 测试并不是越多越好，而是尽可能使你的代码更加健壮。太多测试会增加维护成本，因为每个测试都是需要成本的
		2. Go有内置的计算测试覆盖率的工具，它能帮助你发现没有被测试过的区域。我们不需要追求100%的测试覆盖率，
它只是一个供你获取测试覆盖率的方式。

			运行： go test -cover

 */

func TestSum(t *testing.T) {
/*	numbers := [5]int{1, 2, 3, 4, 5}

	got := integers.Sum(numbers)
	want := 12

	if want != got {
		t.Errorf("got '%d' want '%d' given, %v", got, want, numbers)
	}*/

	// 子测试
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := integers.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := integers.Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}
