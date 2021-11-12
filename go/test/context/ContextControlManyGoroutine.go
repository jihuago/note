package main

import (
	"context"
	"fmt"
	"time"
)

/*
	通过WithCancel() WithDeadline() WithTimeout() WithValue()函数可以创建一棵Context树，树的每个节点都可以有任意多个子节点，节点层级可
有任意多个。
	* WithValue 和取消Context无关，它是为了生成一个绑定键值对数据的Context，这个绑定的数据可以通过Context.Value()方法访问
 */

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "jack")

	go watch(valueCtx, valueCtx.Value(key))
	go watch(valueCtx, "a")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(5 * time.Second)

	// 通知监控停止
	cancel()
	time.Sleep(time.Second)
}

func watch(ctx context.Context, name interface{})  {
	for true {
		select {
		case <- ctx.Done():
			fmt.Println(name, "监控终止.", ctx.Value(key))
			return

		default:
			fmt.Println(name, "监控中...")
			time.Sleep(time.Second)
		}
	}
}


