package dependency_injection

import (
	"fmt"
	"io"
)

/*

	有了注入依赖，我们可以控制数据向哪儿写入，它允许我们：
1. 测试代码
	如果不能轻松地测试函数，这通常是因为有依赖硬链接到了函数或全局状态。DI提倡注入接口，然后就可以在测试中控制你的模拟数据

2. 关注点分离，解耦了数据到达的地方和如何产生数据
如果你感觉一个方法、函数负责太多功能了（比如生成数据并且写入一个数据库？处理HTTP请求并且处理业务级别的逻辑），那么你可能就需要DI这个工具

3. 在不同环境下重用代码
我们的代码所处的第一个环境就是在内部进行测试，但是随后，如果其他人想要用你的代码尝试新的东西，他们只要注入他们自己的依赖即可

 */

func Greet(writer io.Writer, name string) {
	// fmt.Printf会打印到标准输出，用测试框架来捕获它会非常困难
	// 解决：需要做的就是注入打印的依赖，函数不需要关心在哪里打印，如何打印。所以接受一个接口，而非一个具体的类型
	// 如果接受的是一个接口，就可以通过改变接口的实现，控制打印的内容。
	//fmt.Printf("Hello, %s", name)

	fmt.Fprintf(writer, "Hello, %s", name)
}
