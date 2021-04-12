## 编写测试文件
* 文件名： *_test.go
* 测试函数必须以Test开始
* 测试函数只接受一个参数t *testing.T
```html
func TestA(t *testing.T) {}
```

### 运行测试的常用命令

* go test
需要去到测试文件所在的命令
  
* go test -v
go test -v 运行包的测试套件以及示例函数
  
示例函数中需要有： //Output: 6 之类的代码，6是输出的结果

* go test -bench="."
go test -bench="." 来运行基准测试
  
注意：测试文件和源码文件必须位于同一个包内，才可以得到测试代码覆盖率相关数据的
  

### 子测试
在测试函数中，通过`t.Run()`来创建一个子测试。子测试可以进行分组测试，然后再对不同场景进行子测试非常有效


```html

func TestHello(t *testing.T) {
	t.Run("Title", func (t *testing.T) {
        // todo
    )}

    t.Run("Title2", func (t *testing.T){
        // todo
    })
	
}

```

### 编写示例
通过编写示例，让使用者了解你所编写的代码如何使用。

* 示例需要存在一个包的的_test.go文件中的函数
* 示例函数以Example开头
```html
func ExampleAdd() {
    // todo
    // Output: 6
}
```

### 基准测试

* 基准测试格式
```html
func BenchmarkXxx(*testing.B)
```

* 例子
```html
func BenchmarkRepeat(b *testing.B) {
    for i:= 0; i < b.N; i++ {
        Repeat("a")
    }
}

```

* b.N
基准测试运行时，代码会执行b.N次，并测量需要多长时间。
  
测试框架会选择一个它所认为的最佳值，以便获得更合理的结果。
