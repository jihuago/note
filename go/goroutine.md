### goroutine
当你开始一个goroutine时，应该询问两个问题：
1. goroutine什么时候结束
     当你不知道一个goroutine什么时候结束，就不应该启动它
   
2. 如何让它goroutine能够结束

3. 把goroutine的逻辑留给调用者
    
```html

// v1
func main() {
    serve()    
}

// 这种方法不可取，外部调用者不知道serve有goroutine
func serve() {

    go func() {
        // todo serve code
    }()
}

// v2
func main() {
    go serve()
}

func serve() {
    // todo serve code 
}
``` 

