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

#### wrap error
> https://github.com/pkg/errors

通过使用pkg/errors包，可以向错误值添加上下文。

* 如何使用wrap
    * 在你的应用代码中，使用errors.New或errors.Errorf返回错误。
 ```html

    import(
        "github.com/pkg/errors"
    )
    func parseArgs(args []string) error {
        if len(args) < 3 {
            return errors.Error("not enough arguments, expected at least")
        }
    }
```

    * 如果调用其他包内的函数，通常简单的直接返回

```html
if err != nil {
    return err
}
```

在有错误的地方，使用wrap，且不需要记录日志，日记的记录交给上层:
  ```html
errors.Wrap(err, "write failed")
```
 
    * 如果和其他库（github库/基础库）进行协作，考虑使用errors.Wrap或者errors.Wrapf保存堆栈信息。同样适用于和标准库协作的时候。
```html
f, err := os.Open(path)
if err != nil {
    return errors.Wrap(err, "failed to open %q". path)
}
```
  
    * 直接返回错误，而不是每个错误产生的地方到处打日志

    * 在程序的顶部或者是工作的goroutine的顶部（请求入口），使用 %+v把堆栈详情记录
```html
func main() {
    err := app.Run()
    if err != nil {
        fmt.Printf("FATAL: %+v\n", err)
        os.Exit(1)
    }
}
```
    * 使用errors.Cause获取root error，再进行和sentinel error判定

    * 基础库，通用库不选择wrap error
        选择wrap error是只有applications可以选择应用的策略。具有最高可重用性的包只能返回根错误值。

    * 一旦确定函数/方法将处理错误，错误就不再是错误。如果函数/方法依然需要发出返回，则它不能返回错误值。它应该只返回nil

