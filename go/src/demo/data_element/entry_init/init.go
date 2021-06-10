package entry_init

import (
	"data_element/app/controllers"
	"data_element/app/controllers/db"
	"data_element/app/controllers/err"
	"data_element/app/controllers/learnGoWithTest/integers"
	"data_element/common"
	"data_element/goroutime"
	"data_element/router"
)

// 初始化事项
func Init()  {

	r := router.Default()

	r.Get("/str", controllers.Test)
	r.Get("/a", controllers.DemoStr)
	r.Get("/go", goroutime.RunManyGoroutine)
	r.Get("trace", goroutime.DemoTrace)
	r.Get("defer_track", common.DemodeferTrack)
	r.Get("q", controllers.QueryRowDemo)
	r.Get("str1", controllers.DemoStrings)
	r.Get("json", controllers.OmitemptyDemo)
	r.Get("interview", controllers.DemoInterview)
	r.Get("channel", controllers.DemoChannel)
	r.Get("arr", controllers.DemoArr)
	r.Get("debug", controllers.DemoDebug)
	r.Get("new", controllers.DemoAboutNew)
	//r.Get("interface", controllers.DemoGC)
	r.Get("/go-ji-chu", integers.DemoSumAll)
	//r.Get("log", log.LogPrint)
	r.Get("interface", db.DemoGorm)
	r.Get("err", err.DemoErr)

	r.Run()
}
