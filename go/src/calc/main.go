package main

// 引入其它包
import (
	"fmt"
	"os"
	"strconv"
	"calc/simplemath"
	"calc/funcs"
	"sync"
)

// 定义一个用于打印程序使用指南的函数
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}

var wg sync.WaitGroup

// 程序入口函数
func main ()  {
	/*
	 * 用于获取命令行参数，注意程序名本身是第一个参数，
	 * 比如 calc add 1 2 这条指令，第一个参数是 calc
	 */
	args := os.Args
	// 除程序名本身外，至少需要传入两个其它参数，否则退出
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	// 第二个参数表示计算方法
	switch args[1] {
	// 如果是加法的话
	case "add":
		// 至少需要包含四个参数
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 获取待相加的数值，并将类型转化为整型
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		// 获取参数出错，则退出
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 从 simplemath 包引入 Add 方法进行加法计算
		ret := simplemath.Add(v1, v2)
		// 打印计算结果
		fmt.Println("Result: ", ret)
	// 如果是计算平方根的话
	case "sqrt":
		// 至少需要包含三个参数
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 获取待计算平方根的数值，并将类型转化为整型
		v, err := strconv.Atoi(args[2])
		// 获取参数出错，则退出
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 从 simplemath 包引入 Sqrt 方法进行平方根计算
		ret := simplemath.Sqrt(v)
		// 打印计算结果
		fmt.Println("Result: ", ret)
	// 如果计算方法不支持，打印程序使用指南
	case "down":
		for i := 0; i < 3; i++ {
			wg.Add(1)
			// go funcs.Download()启动新的协程并发执行download函数
			go funcs.Download("www.baidu.com/" + string(i+'0'))
		}
		wg.Wait()
		fmt.Println("Done!")
		break
	default:
		Usage()
	}
}