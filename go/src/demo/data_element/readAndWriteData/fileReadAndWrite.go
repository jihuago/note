package readAndWriteData

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件读写
// 在Go语言中，文件使用指向os.File类型的指针来表示的，也叫做文件句柄。
func ReadFile()  {
	inputFile, err := os.Open("./readAndWriteData/input.txt")

	if err != nil {
		fmt.Printf("open fail")
		return
	}

	defer inputFile.Close()

	var users [5]struct{}
	fmt.Printf("%l", users)
}

func ReadTxtFile() {
	inputFile, err := os.Open("./public/a.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)

	for {
		content, err := reader.ReadString('\n')
		fmt.Printf("The input was: %s", content)

		if err == io.EOF {
			fmt.Println(err)
			return
		}

	}

}

// 带缓冲的读取
// 如果文件的内容是不按划分的，或者干脆就是一个二进制文件。可以使用bufio.Reader的Read()

// 按列读取文件中的数据
// 如果数据是按列排列并用空格分隔的，你可以使用fmt包提供的以FScan开头的一系列函数来读取他们
func ReadFile2()  {
/*	file, err := os.Open("./public/product2.txt")

	if err != nil {
		panic(err)
	}*/
}
