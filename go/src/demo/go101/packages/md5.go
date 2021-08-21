package packages

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5Demo()  {
	h := md5.New()
	io.WriteString(h, "123456")
	// 一般128位的MD5散列被表示为32位16进制数字，所以是%x
	fmt.Printf("%x\n", h.Sum(nil))

	bt := []byte{'A', 'B', 'a', 1, '*'}
	fmt.Printf("%s\n", bt)
	fmt.Printf("%x\n", bt)
	//fmt.Printf("%T", bt)
	fmt.Printf("%+v\n", bt)
	fmt.Println(bt)
	fmt.Printf("%d \n", 6 << 1)
}
