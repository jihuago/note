package main

import "data_element/entry_init"

func main()  {
	// 数据类型

/*	// 整型
	var number1 int16
	number1 = 12

	var number2 int8
	number2 = 12

	// 不同类型的值不能放一起比较
	//res1 := (number2 == number1)

	// 不同类型的整型值不能直接进行算数运算 int8()转换类型为int8
	res := number2 + int8(number1)
	res ++

	// 各种类型的整型变量都可以直接与字面常量进行比较
	res4 := res == 2

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(res)
	fmt.Println(res4)*/

	//firstChr, length := common.GetStringFirstChar("啊test")
	//fmt.Println(firstChr)
	//fmt.Println(length)
	//
	//str := "啊test"
	//fmt.Println(len(str))
	//
	//common.TestDataTranfer()
	//
	//common.ArrMethod()

	// 字典的使用
	//common.MapTest()
	//common.TestPointer()

	// 控制结构
	//fmt.Println(common.RelaxTip())

	//函数
	//common.Run()

	// 正则
	//common.Test()

	// 结构体
	//common.TestOuterS()
	//common.TestA()

	// 文件操作
	//readAndWriteData.GetUserInput()
	//readAndWriteData.SwitchInput()
	//readAndWriteData.ReadTxtFile()
	//readAndWriteData.ReadFile2()
	//readAndWriteData.ReadCsv()
	//readAndWriteData.WriteFile()
/*	page := readAndWriteData.Page{"./public/page.txt", []byte{
		'a', 'b',
	}}

	//page.Save()
	c, err := page.Load("./public/page.txt")
	fmt.Println(c, err)*/

	// 文件拷贝
	//i, err := readAndWriteData.CopyFile("./public/dst1.txt", "./public/a2.txt")
	//fmt.Println(i, err)

	// 从命令韩读取参数
	//common.TestArgs()
	//common.HelloWho()

	// panic
	//common.Close()
	//common.TestRecover()
	//common.Do()

	// 类型的String()方法
	//p := common.Person1{}
	//p.Init()

	// 垃圾回收和SetFinalizer
	//common.GetMemStatus()

	// 接口与反射
	//_interface.Init()

/*	var s _interface.Simple
	_interface.InitSimpler(&s)

	// 类型断言
	_interface.CheckVarI()

	_interface.TestEmptyInterface()*/

	// 接口，写文件
/*	var arr []string
	arr = []string{
		"this", "name",
	}
	fileLog := log.FileLog{"./public/2021.log", arr}
	fileLog.Write()*/


	//reflect.TestReflect()

	// 协程
	//goroutime.Testgoroutime()
	//goroutime.TestChannle()
	//goroutime.Testf1()
	//goroutime.TestChannelBuf()
	//goroutime.TestGoFor()
	//goroutime.Test()

	// fmt包
	//common_package.TestFmt()
	//common_package.TestTime()
	//common_package.DemoTicker()

	// http
	//common_package.DemoGet()
	//common_package.DemoTestWithParam()

	entry_init.Init()
}
