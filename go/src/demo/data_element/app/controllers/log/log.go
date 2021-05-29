package log

import (
	"fmt"
	"log"
	"os"
)

func LogPrint()  {

	logFile, err := os.OpenFile("./runtime/logs/test.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.SetPrefix("俏手艺：")

	// logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)
	log.Println("log1")
	log.Printf("this is a log: %s", "hah")

	// 标准logger的配置
	// 默认情况下的logger只会提供日志的时间信息，但是很多 情况想要得到更多信息，比如记录日志的文件名和行号等
	// log标准库为我们提供了定制这些设置的方法
	/*
		log标准库中的Flags函数会返回标准logger的输出配置，
		而SetFlags函数用来设置标准logger的输出配置

	flag选项
		log标准库提供了如下flag选项，他们是一些定义好的常量

		const (
			//
		)

	配置日志输出位置
		func SetOutput(w io.Writer)
		SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出
	 */











}
