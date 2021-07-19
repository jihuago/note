package slice

/*
	* slice 即动态数组，其长度并不固定，可以向切片中追加元素，它会在容量不足时自动扩容
	* slice 声明
		t := []int
	* slice数据结构
		type SliceHeader struct {
			Data uintptr // 指向数组的指针
			Len int // 当前切片的长度
			Cap int // 当前切片的容量
		}
	slice扩容方法：
		1. 如果期望容量大于当前容量的两倍就会使用期望容量
		2. 如果当前切片的长度小于1024就会将容量翻倍
		3. 如果当前切片的长度大于1024就会每次增加25%的容量，直到新容量大于期望容量
 */
func DeepSlice()  {

}
