package main

import "fmt"

func main() {

	// 数组是长度固定的数据类型，必须存储一段相同类型的元素，而且这些元素是连续的
	// Go语言规定，必须是长度一样，并且每个元素的类型也一样的数组，才是同样类型的数组

	// 声明数组
	var names [5]string
	fmt.Println(names) // 未初始化，每个元素都是对应类型的零值，string就是空字符串

	names = [5]string{"jack", "mary"} // 初始化数组
	fmt.Println(names)

	// 常用的方式
	ages := [5]int{18, 20}
	fmt.Println(ages)

	// 不指定数组长度，Go自动推导
	infos := [...]float32{18.2, 2: 10}
	fmt.Println(infos)

	for k, info := range infos {
		fmt.Printf("索引：%d 值：%.2f\n", k, info)
	}

	// 指针数组
	address := [3]*int{1: new(int)}
	// 给索引1赋值
	*address[1] = 2

	// 因为只有索引1创建了内存空间，所以只能给索引1赋值，如果给索引0赋值，就会提示无效内存
	//*address[0] = 1 // panic: runtime error: invalid memory address or nil pointer dereference

	fmt.Println(ages)
	modify(&ages)
	fmt.Println(ages)

	//定义多维数组
	employees := [2][2]int{
		{1, 1},
		{2, 2},
	}
	fmt.Println(employees)

	a1 := [...]int{1, 2, 3, 4}
	result := SumArray(a1)
	fmt.Println(result)

	testArr := []int{1, 3, 5, 7, 8}
	keys := TwoNumberSum(testArr, 89)
	fmt.Println("结果:", keys)

	t := twoSum(testArr,8)
	fmt.Println(t)

}

// 函数间传递数组  *[5]int 数组指针
func modify(a *[5]int)  {
	a[1] = 29
}

// SumArray 求某个数组内所有元素的和   求数组[1, 3, 5, 7, 8]所有元素的和
func SumArray(arr [4]int) (result int) {
	result = 0
	for _, val := range arr {
		result += val
	}

	return
}

// TwoNumberSum 数组中两数之和
// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func TwoNumberSum(nums []int, target int) [][]int  {
	m := make(map[int]int)
	result := make([][]int, 0)
	for k, num := range nums {
		diff := target - num
		if i, ok := m[diff]; ok {
			result = append(result, []int{i, k})
		} else {
			m[num] = k
		}
	}

	return result
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, _ := range nums {
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			return []int{i, j}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}
