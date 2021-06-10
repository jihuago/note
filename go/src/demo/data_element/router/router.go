package router

import (
	"fmt"
	"os"
)

type route struct {
	routrMap map[string]func()
	run func()
	uri string
}

type router interface {
	Get(pathname string, fun func())
	Run()
}

func (r *route) Get(uri string, fun func()) {

	if re := recover(); re != nil {
		defer func() {
			fmt.Println("not found:", re)
		}()
	}

	// 根据命令行传递的参数执行对应代码
	args := os.Args
	pathName := args[1]

	_, ok := r.routrMap[uri]
	if !ok {
		r.routrMap[uri] = fun
	}

	r.uri = pathName
	r.run = r.routrMap[pathName]
	//fmt.Println()
}

func Default() *route  {
	r := &route{}
	r.routrMap = make(map[string]func())
	return r
}

func (r *route) Run(addr ...string)  {

	defer func() {
		if re := recover(); re != nil {

			err := fmt.Errorf("Error not found uri: %v, detail: %v", r.uri, re)
			fmt.Println(err)
		}
	}()

	run := r.run
	run()
}
