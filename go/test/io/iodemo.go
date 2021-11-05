package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// 在io包中最重要的是两个接口，Reader和Writer接口。

	/*

		Read()将len(p)个字节读取到p中。它返回读取的字节数n(0<=n<=len(p))以及以及任何遇到的错误
		如果可读取的数据不到len(p)个字节，Read会返回可用数据，而不是等待更多数据
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	 */

/*	// 从文件中读取
	file, err := os.Open("./ab.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	for true {

		data, err := ReadFrom(file, 1)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("读取失败", err)
			return
		}
		fmt.Printf("%s", data)
	}

	f, err := os.OpenFile("./a.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	writeErr := Write(f, "this is a test")
	if writeErr != nil {
		fmt.Println("错误：", writeErr)
	} else {
		fmt.Println("写入成功")
	}*/

	file1, err1 := os.Open("./a.txt")

	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer file1.Close()

	contenxt ,_ := ioutil.ReadAll(file1)
	fmt.Println(string(contenxt))
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}
