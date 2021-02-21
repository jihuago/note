package goroutime

import (
	"fmt"
	"os"
	"runtime/trace"
)

// trace记录了运行时的信息，能提供可视化的Web界面
func DemoTrace() {
	// 创建trace文件
	f, err := os.Create("./public/trace.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Hello trace")
}
