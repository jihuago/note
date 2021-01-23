package readAndWriteData

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 写文件

func WriteFile()  {
	fileHandler, err := os.OpenFile("./public/test.txt", os.O_WRONLY|os.O_CREATE, 066)

	if err != nil {
		panic(err)
	}
	defer fileHandler.Close()

	// 不使用缓冲区，直接将内容写入文件
	fileHandler.WriteString("test\n")
}

/*
	编写一个save方法，将Title作为文件名，Body作为文件内容，写入到文本文件中
*/
type Page struct {
	Title string
	Body []byte
}

func (this *Page) Save()  {
	file, err := os.OpenFile(this.Title, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range this.Body {
		file.WriteString(string(v))
	}

}

/*
	读取文件内容
 */
func (this *Page) Load(title string) (content string , err error)  {

	// 将整个文件的内容读到一个字符串
	buf, err := ioutil.ReadFile(title)
	if err != nil {
		panic(err.Error())
	}

	return string(buf), err

}



