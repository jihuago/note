package entry_init

import (
	"data_element/app/controllers"
	"data_element/router"
)

// 初始化事项
func Init()  {

	r := router.Default()

	r.Get("/str", controllers.DemoStr)
	r.Get("/a", controllers.DemoStr)

	r.Run()
	//fmt.Println(r)
}
