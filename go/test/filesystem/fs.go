package main

import "fmt"

func main() {

	// io/fs的加入让我们易于面向接口编程，而不是面向os.File这个具体实现。

	var str string = "test"
	var data []byte = []byte(str)

	f := &file{}
	res, err := f.Read(data)
	fmt.Println(res, err)

}


