package integers

func DemoSumAll()  {
	SumAll([]int{1, 2}, []int{5, 6})
}

/*// 可变参数  ...类型
func SumAll(numbersToSum ...[]int) (sum []int) {
	//return []int{3, 9}

	lengthOfNumbers := len(numbersToSum)

	// 根据传递进来的参数个数，创建切片
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}*/

func SumAll(numbersToSum ...[]int) (sum []int) {

	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int  {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
