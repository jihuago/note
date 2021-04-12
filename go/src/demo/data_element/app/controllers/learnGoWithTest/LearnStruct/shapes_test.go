package LearnStruct

import "testing"

// 需求：计算一个给定高和宽的长方形的周长

// 1. 先写测试函数
func TestPerimeter(t *testing.T)  {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %f", got, want)
	}
}

// 2. 运行测试

// 3. 为运行测试函数编写最少的代码并检查失败时的输出
func Perimeter(width, height float64) float64 {
	return 0
}
