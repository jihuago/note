package common

import (
	"fmt"
	"sort"
)

// 字典声明 testMap是字典变量名  string 是键的类型，int是存放的值类型
var testMap map[string]int
func MapTest()  {

	// 赋值
	testMap = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	fmt.Println(testMap)

	// 声明字典的另外一种方式，通过这种方式初始后，方可直接往字典添加键值对
	var info = make(map[string]string)
	info["name"] = "jack"
	info["sex"] = "mary"

	// 创建时指定字典的初始存储能力（超出会自动扩容）
	//var names = make(map[string]string, 10)

	fmt.Println(info)
	// 查找元素
	value, ok := info["name"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found jack in dict info")
	}

	// 删除字典元素
	delete(info, "sex")

	// 遍历字典
	for key, value := range testMap {
		fmt.Println(key, value)
	}

/*	for _, value := range testMap {
		fmt.Println(value)
	}

	// 只获取到Key
	for key := range testMap {
		fmt.Println(key)
	}*/

	// 键值对调
	// 键值对调就是交换字典的键和值
	fmt.Println("\r\n键值对调")

	userInfo := map[string]string {
		"name" : "jack",
		"sex" : "男",
		"height": "176cm",
	}
	changeInfo := make(map[string]string)
	for k, v := range userInfo{
		changeInfo[v] = k
	}

	fmt.Println(changeInfo)

	// 字典排序
	// 如果想要对字典排序，可以通过分别对字典的键和值创建切片，然后通过对切片进行排序来实现
	fmt.Println("\r\n字典排序")
	classInfo := map[string]string {
		"A" : "jack",
		"C" : "jack",
		"B" : "mary",
	}

	keys := make([]string, 0)

	for k, _ := range classInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	fmt.Println(classInfo)



}