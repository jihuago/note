package main

import (
	"gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(context *gee.Context) {
		
	})

	r.Run(":9999")
}
