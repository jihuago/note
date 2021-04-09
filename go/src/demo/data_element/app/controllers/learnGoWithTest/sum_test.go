package learnGoWithTest

import (
	"data_element/app/controllers/learnGoWithTest/integers"
	"testing"
)

/*
		numbers := [5]int{1, 2, 3, 4, 5}
		numbers := [...]int{1, 2, 3}
 */

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := integers.Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got '%d' want '%d' given, %v", got, want, numbers)
	}
}
