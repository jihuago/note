package _map

import "fmt"

// map的长度是不固定的，也就是和slice一样，也是一种引用类型
// map和其他基本类型不同，map不是thread-safe(线程安全)，在多个goroutine存取时，必须使用mutex lock机制
func MapDemo()  {
	//var numbers map[string]int

	info := make(map[string]string)
	info["name"] = "jack"
	info["sex"] = "女"

	fmt.Println(info)

	rating := map[string]float32{"C": 5, "GO": 4.5, "Python": 4.5}
	ra, ok := rating["GO"] // 判断key是否存在
	if ok {
		fmt.Println("GO键存在", ra)
	}



}
