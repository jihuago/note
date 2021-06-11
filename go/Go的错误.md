## 处理错误的基本策略
* 哨兵错误
    哨兵错误的处理方式通过特定值表示成功和不同错误，依靠调用方对错误进行检查：
  ```html
        if err === Errsomething {...}
```
例如，io.EOF = errors.New("EOF")
这种错误处理的方式引入了上下层代码的依赖，如果被调用方的错误类型发生了变化，则调用方也需要对代码进行修改

* 隐式错误
> if err != nil {return err}

这种错误处理的方式直接返回错误的任何细节，直接将错误进一步报告给上层。这种情况下，错误在当前调用方这里完全没有进行任何加工，与没有进行处理几乎是等价的
这回产生一个致命的问题在于：丢失调用的上下文信息，如果某个错误连续向上层传播了多次，那么上层代码可能在输出某个错误时，根本无法判断该错误的错误信息从
哪儿传播出来。
```

尽量少用哨兵错误

* 尽量少用panic
    Go panic意味着fatal error(程序挂了)，不能假设调用者来解决panic，业务逻辑不用panic；
  
对于真正意外的情况，哪些表示不可恢复的程序错误，例如索引越界、不可恢复的环境问题、栈溢出，才用panic


* Error Types（尽量避免使用）

Error Type 是实现了error接口的自定义类型。例如MyError类型记录了文件和行号以展示发生了什么。

```html
type MyError struct {
    Msg string
    File string
    Line int
}
func (e *MyError) Error() string {
    return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}
func test() error {
    return &MyError{"Something happend", "server.go", 34}
}

func main() {
    err := test()
    // Error Type要使用类型断言和类型switch
    switch err := err.(type) {
        case nil:

        case *MyError:
            fmt.Println("error occured on line:", err.Line)
        default:
    }
}
```
Error Type要使用类型断言和类型switch，就要让自定义的error变为public。这种模型会导致和调用者产生强耦合
，从而导致API变得脆弱。
建议：避免错误类型，至少避免将他们作为公共API的一部分。

* Opaque error
Opaque error不透明的错误处理策略。这是最灵活的错误处理策略，因为它要求代码和调用者之间的耦合最少。
  不透明错误处理：只需返回错误不假设其内容。
  
```html
func fn() error {
    x, err := bar.Foo()
    if err != nil {
        return err
    }
}
```

#### 错误处理
* 保证只处理错误一次
* 错误要被日志记录
* 应用程序处理错误，保证100%完整性
* 之后不再报告当前错误