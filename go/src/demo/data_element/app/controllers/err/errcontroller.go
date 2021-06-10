package err

import (
	"errors"
	"fmt"
	"time"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func NewError(text string) error {
	return errorString{text}
}

var ErrType = NewError("EOF")

type FileError struct {
	
}

func (fe *FileError) Error() string {
	return "文件错误"
}

func DemoErr()  {
	if ErrType == NewError("EOF") {
		fmt.Println("Error:", ErrType)
	}

	if ErrType == errors.New("EOF") {
		fmt.Println("EOF")
	}

	fe := FileError{}
	if "文件错误" == fe.Error() {
		fmt.Println(fe.Error())
	}

	// 使用自封装的goroutine
	Go(func() {
		fmt.Println("hello")
		panic("不知道晚为什么panic")
	})
	time.Sleep(5 * time.Second)
}

// 封装一下goroutine，捕获panic，防止使用原生的go func () {}
func Go(x func())  {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// 这里处理一下panic() 避免有人在goroutine使用了panic()导致整个进程挂掉
				fmt.Println(err)
			}
		}()

		x()
	}()
}