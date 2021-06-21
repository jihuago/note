package err

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

// github.com/pkg/errors的例子
func PkgErr()  {
/*	err := test1()

	switch errStr := errors.Cause(err).(type) {
	case error:
		fmt.Println("文件读取错误", errStr)
		break
	default:
		fmt.Println("nothing")
	}*/

	/*

	因为stack.go有StackTrace()方法
	func (s *stack) StackTrace() StackTrace {
		f := make([]Frame, len(*s))
		for i := 0; i < len(f); i++ {
			f[i] = Frame((*s)[i])
		}
		return f
	}

	https://pkg.go.dev/github.com/pkg/errors#StackTrace

	*/
	type stacker interface {
		StackTrace() errors.StackTrace
	}

	// 如果转换合法，err是stackTracer类型的值
	err, ok := errors.Cause(fn()).(stacker)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v\n", st[0:2])

/*	e1 := fn1()
	fmt.Printf("%+v\n", e1)*/

	fnAssertType()

}

func test1() error {
	// errors.Wrap函数返回一个新错误，该错误将上下文添加到原始错误。

	//err := fmt.Errorf("this is err test")

	f, err := os.OpenFile("./2.txt", os.O_RDONLY, 0644)
	defer f.Close()

	return errors.Wrap(err, "read failed")
}

func fn() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")

	return errors.Wrap(e3, "outer")
}

func fn1() error {
	e1 := errors.New("error")
	err := errors.WithStack(e1)

	return err
}

type tester interface {
	test() string
}

type demo struct {
}

func (d *demo) test() string {
	return "this is a test"
}

type demo1 struct {

}

// 类型断言例子
func fnAssertType()  {

	var t tester

	d := new(demo)
	t = d

	if v, ok := t.(*demo); ok {
		fmt.Println(v)
	}

	getTester().(*demo)

}

func getTester() tester {
	return &demo{}
}

