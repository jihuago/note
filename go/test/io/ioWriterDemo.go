package main

import (
	"fmt"
	"io"
)

/*

	type Writer interface {

		// Write()方法将len(p)个字节从p中写入到基本数据流中。
		Write(p []byte) (n int, err error)
	}
 */
func Write(writer io.Writer, str string) error {
	p := make([]byte, 1)
	p = []byte(str)
	_, err := writer.Write(p)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
