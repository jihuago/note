package generator

import "fmt"

// 生成器是指当被调用时返回一个序列中下一个值的函数
// 生成器每次返回的是序列中下一个值而非整个序列；这种特性也叫惰性求值

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0

	go func() {
		for true {
			yield <- count
			count ++
		}
	}()

	return yield
}

func generateInteger() int {
	return <- resume
}

func DoGenerate()  {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}

