package readAndWriteData

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	* 如何读取用户的键盘输入：
		1. 从键盘和标准输入os.Stdin读取输入，最简单的办法是使用fmt包提供的Scan和Sscan开头的函数
		2. Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行
		3. Scanf与其类似，除了Scanf的第一个参数用作格式字符串，用来决定如何读取
	* 使用bufio包提供的缓冲读取来读取数据
*/

func GetUserInput() {
	var (
		//name, gender, s string
		s string
		i int
		f float32
		input = "56.12 / 5212 / Go"
		format = "%f / %d / %s"
	)

/*	fmt.Println("Please enter your name and gender:")
	fmt.Scanln(&name, &gender)
	fmt.Printf("Hi %s %s\n", name, gender)*/

	// 获取用户输入的值
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)

	var (
		inputReader *bufio.Reader
		input2 string
		err error
	)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")

	// ReadString(delim byte) 该方法从输入中读取内容，直到碰到delim指定的字符，然后将读取到的内容连同delim字符串一起放到缓冲区
	// ReadString返回读取到的字符串，如果碰到错误则返回nil。
	input2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was : %s\n", input2)
	}
}

var nrchars, nrwords, nrlines int

func SwitchInput() {
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter some input, type S to stop：")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}

		if input == "S\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters:%d\n")
		}

		counters(input)
	}

}

func counters(input string)  {
	nrchars += len(input) - 2
	nrwords += len(strings.Fields(input))
	nrlines ++
}