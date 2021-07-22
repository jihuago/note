package slice

import "fmt"

// slice是一个引用类型，slice总是指向一个底层数组
func SliceDemo()  {
	//sliceDefine()
	sliceOp()
}

func sliceDefine() {
	// slice 切片声明
	var fslice []int
	fslice = []int{1, 2} // 赋值

	// 声明并初始化数据slice
	byteSli := []byte{'a', 'b', 'c'}
	fmt.Println(byteSli, fslice)
}

// slice可以从一个数组或一个已经存在的slice中再次声明
// slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i
func sliceOp()  {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f'}
	a := ar[3:5]
	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := ar[2:4:4]
	fmt.Println(string(b))
}


