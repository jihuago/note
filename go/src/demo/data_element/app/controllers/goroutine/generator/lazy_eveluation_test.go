package generator

import (
	"testing"
)

const maxNum = 100

func TestDoGenerate(t *testing.T) {
/*	t.Run("test for Generate", func(t *testing.T) {
		DoGenerate()
	})*/

/*	t.Run("test for future", func(t *testing.T) {
		got:= Sum([]int{1, 2, 3, 4})
		want := 10

		if want != got {
			t.Errorf("want %d, but got %d", want, got)
		}

	})*/
}

func BenchmarkSum(b *testing.B) {

	testArr := []int{}

	for i :=0; i < maxNum; i++ {
		testArr = append(testArr, i)
	}
	//fmt.Println(testArr)
	for i := 0; i < b.N; i++ {
		Sum(testArr)
	}
}

func BenchmarkSumNogoroutine(b *testing.B) {

	testArr := []int{}

	for i :=0; i < maxNum; i++ {
		testArr = append(testArr, i)
	}

	for i := 0; i < b.N; i++ {
		SumNoGoroutine(testArr)
	}
}