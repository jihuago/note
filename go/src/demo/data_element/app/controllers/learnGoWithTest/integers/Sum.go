package integers

// go数组的大小也属于类型的一部分。
func Sum(numbers []int) (sum int)  {
	sum = 0

/*	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}*/

	for _, number := range numbers {
		sum += number
	}

	return
}
