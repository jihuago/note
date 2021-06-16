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
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	err, ok := errors.Cause(fn()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v\n", st[0:2])


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

