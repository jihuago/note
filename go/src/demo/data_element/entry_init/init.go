package entry_init

import (
	"data_element/app/controllers"
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

	r.Run()
}
