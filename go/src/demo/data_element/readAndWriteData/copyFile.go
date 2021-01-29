package readAndWriteData

import (
	"io"
	"os"
)

/*
	文件拷贝
		1. 使用io包
*/
func CopyFile(dstName, srcName string) (written int64, err error)  {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	defer src.Close()

	dst, err := os.Create(dstName)

	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

