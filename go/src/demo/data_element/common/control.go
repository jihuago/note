package common

import (
	"time"
)

// 控制结构 if else  switch
// 休息日提醒
func RelaxTip() bool {

	relaxDay := getRelaxDate()
	currentDay := int(time.Now().Weekday())

	testFor()

	for _, value := range relaxDay {
		if value == currentDay {
			return true
		}
	}


	return false
}

func getRelaxDate() map[string]int {
	return map[string]int {
		"Saturday": 6,
		"Sunday": 7,
	}
}

func testFor()  {
	// 打印0-10的数
/*	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}*/

	/*
		基于条件判断的迭代
			for 条件语句 {}


	var i int = 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	*/

	/*
			*****
			*****
			*****
	       宽 5 高 3


	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	 */

	/*
		无限循环
			for {}

		for-range结构
			for index, val := range coll {}
			val 为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所作的任何修改都不会影响（
				如果val为指针，则会产生指针的拷贝，就可以修改集合整的原值
			）
	*/



}
