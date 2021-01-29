package log

import (
	"fmt"
	"os"
)

// 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。
type Loger interface {
	Write()
}

type FileLog struct {
	Filename string
	LogInfo interface{}
}

type Console struct {
	logInfo interface{}
}

func (f *FileLog) Write()  {
	file, err := os.OpenFile(f.Filename, os.O_CREATE|os.O_APPEND, 0600)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	str := fmt.Sprintf("Log: %v \n", f.LogInfo)

	file.WriteString(str)
}
