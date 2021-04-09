package integers

// 可变参数  ...类型
func SumAll(numbersToSum ...[]int) []int {
	//return []int{3, 9}

	lengthOfNumbers := len(numbersToSum)

	// 根据传递进来的参数个数，创建切片
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
