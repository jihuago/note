package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// 空的context，
	process(ctx)

	// WithValue() 函数能够将请求作用域的数据与Context对象建立关系
	// 仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数
	ctx = context.WithValue(ctx, "traceId", "qcrao-2020")
	process(ctx)
}

func process(ctx context.Context)  {
	// Value()方法会从Context中返回键对应的值
	traceId, ok := ctx.Value("traceId").(string)

	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}
