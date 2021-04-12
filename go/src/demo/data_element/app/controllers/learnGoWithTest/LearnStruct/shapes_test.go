package LearnStruct

import "testing"

// 需求：计算一个给定高和宽的长方形的周长

// 1. 先写测试函数
func TestPerimeter(t *testing.T)  {

	// version 1
/*	// 表格驱动测试
	t.Run("test Circle", func(t *testing.T) {
		areaTests := []struct{
			shape Shapes
			want float64
		}{
			{&Rectangle{12, 6}, 72.0},
			{&Circle{10}, 314.1592653589793},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		}

	})*/

	areaTests := []struct{
		name string
		shape Shapes
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: &Circle{12}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		// go test -run TestPerimeter/Rectangle
		// 通过上述命令来运行列表中指定的测试用例
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}

}

// 2. 运行测试

// 3. 为运行测试函数编写最少的代码并检查失败时的输出
func Perimeter(r *Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
