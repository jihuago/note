package generator

// 所谓Futures就是指：有时候在你使用某一个值之前需要先对其进行计算。这种情况下，你就可以在另一个处理器上进行该值的计算，到使用时，该值就已经计算完毕

// 计算切片的总和：利用Futures模式。把切片分为两部分，分别计算，最后加起来
func sum(slice []int) chan int {
	future := make(chan int, 2)


	go func() {
		future <- doSum(slice)
	}()

	return future
}

func doSum(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func Sum(arr []int) int  {

	arrLen := len(arr)

	ch1 := sum(arr[0:arrLen / 2])
	ch2 := sum(arr[arrLen / 2 : ])

	res1 := <- ch1
	res2 := <- ch2

	return res1 + res2
}

func SumNoGoroutine(arr []int) int {

	sum := doSum(arr)

	return sum
}