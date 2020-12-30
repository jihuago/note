package simplemath
// 在编写单元测试，需要引入testing包，可以基于该包提供的方法来实现自动化测试，测试方法的格式如下：
/*
		func TestXXX(t *testing.T) {
			// 测试逻辑
		}

		// 运行单元测试
		$ go test simplemath
 */
import "testing"

func TestAdd(t *testing.T)  {
	r := Add(1, 2)
	if r != 3 {
		t.Error("Add(1, 2) failed. Got %d, expected 3", r)
	}
}