package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// path/filepath包：兼容操作系统的文件路径操作

	// 1. 路径分隔符：os.PathSeparator，下面列子拼接出 Linux下： /main  Window: \main
	var build = strings.Builder{}
	build.WriteString(string(os.PathSeparator)) // 拼接字符串的方式，效率较高
	build.WriteString("main")
	res := build.String()
	fmt.Println(res)

	// 2. 解析路径名字符串
	// Dir() 和 Base() 函数将一个路径名字符串分解成目录和文件名梁
	pathStr := "/home/linz/index"
	baseName := filepath.Base(pathStr)
	dirName := filepath.Dir(pathStr)

	// 3. Clean() 规整路径
	cleanStr := filepath.Clean("/./root/a/./")

	// 4. 路径的切分和拼接
	// Split() 函数根据最后一个路径分隔符将路径分隔为目录和文件名两部分
	filepath.Split("/home/linz/studygolang")

	fmt.Println(baseName, dirName, cleanStr)
}
