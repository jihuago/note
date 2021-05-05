package goroutine

import "fmt"

// 通道可以被显式的关闭，尽管它们和文件不同：不必每次都关闭。
func RunGoroutine3()  {
	ch := make(chan string)
	go sendData1(ch)
	getData1(ch)
}

func getData1(ch chan string) {
	for true {
		input, open :=<- ch

		// !open 检查通道是否还有内容
		if !open {
			break;
		}

		fmt.Println(input)
	}

	/*
		使用for- range语句来读取通道是更好的方法，因为这会自动检测通道是否关闭

	 */
}

func sendData1(ch chan string)  {
	ch <- "A"
	ch <- "B"

	// 关闭通道
	defer close(ch)
}