package testing

import (
	"data_element/strev"
	"testing"
)

/*
	* 对一个包做单元测试，需要写一些可以频繁执行的小块测试单元来检查代码的正确性
	* 测试程序必须属于被测试的包，并且文件名满足*_test.go，测试代码和业务代码是分开的
	* _test程序不会被普通的Go编译器编译，所以当放应用部署到生产环境时它们不会被部署，只有gotest会编译所有的程序：普通程序和测试程序
	* 测试文件中必须导入"testing"包，并写一些名字已TestZzz开头的全局函数，这里Zzz是被测试函数的字母描述
	* 测试函数必须有这种形式的头部
		func TestAbcde(t *testing.T)

		T是传给测试函数的结构类型，用来管理测试状态，支持格式化测试日志，如t.Log, t.Error, t.ErrorF
		在函数的结尾把输出跟想要的结果对比，如果不等就打印一个错误

	* 用下面这些函数来通知测试失败
		1. func (t *T) Fail()
			标记测试函数为失败，然后继续执行
		2. func (t *T) FailNow()
			标记测试函数为失败并终止执行；文件中别的测试页被略过，继续执行下一个文件
		3. func (t *T) Log(args ...interface{})
			args 被用默认的格式格式化并打印到错误日志中
		4. func (t *T) Fatal(args ...interface{})

	* 运行go test来编译测试程序，并执行程序中所有的TestZzz函数。如果所有的测试都通过会打印PASS

	* gotest可以接受一个或多个函数程序作为参数，并指定一些选项
		结合 --chatty 或 -v 选项，每个执行的测试函数以及测试状态会被打印
		go test fmt_test.go --chatty

	* testing包中有一些类型和函数可以用来做简单的基准测试；测试代码中必须包含已BenchmarkZzz开头的函数并接受一个 *testing.B类型的参数
		func BenchmarkReverse(b *testing.B) {}
	* 命令go test -test.bench=.*会运行所有的基准测试函数；
*/

type ReverseTest struct {
	in, out string
}

var ReverseTests = []ReverseTest{
	{"ABCD", "DCBA"},
	{"CVO-AZ", "ZA-OVC"},
	{"Hello 世界", "界世 olleH"},
}

func TestReverse(t *testing.T)  {
	// testing with a battery of testdata:
	for _, r := range ReverseTests {
		exp := strev.Reverse(r.in)

		if r.out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", r.in, exp, r.out)
		}
	}
}

func BenchmarkReverse(b *testing.B)  {
	s := "ABCD"
	for i := 0; i < b.N; i++ {
		strev.Reverse(s)
	}
}

