package common

// 结构体
// 定义
type info struct {
	name string
	age int
}

/*

结构体定义的一般方式：
	type identifier struct {
		field1 type1
		field2 type2
	}
	结构体里的字段都有名字，如field1，如果字段在代码中从来也不会被用到，那么可以命名为_

	结构体的字段可以是任何类型，甚至是结构体本身，也可以是函数或者接口。

	赋值
		type T struct {a, b int}
		var s T
		s.a = 5
		s.b = 8

	使用new
		使用new函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)

		惯用方法：t := new(T)
		声明var t T 也会给t分配内存，并零值化内存。t通常被称为类型T的一个实例或对象
	结构体的内存布局
		Go语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

	递归结构体
		结构体类型可以通过引用自身来定义。这在定义链表或二叉树的元素时特别有用，此时节点包含指向临近节点的链接。

 */

// 定义一个链表： head su => data su => tail nil
// data字段用于存放有效数据，su指针指向后继节点
// 链表中的第一个元素叫head，它指向第二个元素；最后一个元素叫tail，它没有后继元素，所以su为nil值。
type Node struct {
	data float64
	su *Node
}

// 定义一个双向链表节点
type DoubleNode struct {
	prev *DoubleNode
	data []string
	next *DoubleNode
}

func Person(p *info)  {

	// 赋值
	//p1 := info{"jack", 23}

}

/*
	使用工厂方法创建结构体实例
		* Go语言不支持面向对象编程语言那样的构造子方法，但是可以会很容易在GO中实现。
			type File struct {
				fd int // 文件描述符
				name string // 文件名
			}
			// 下面是这个结构体类型对应的工厂方法，它返回一个指向结构体实例的指针
			func NewFile(fd int, name string) *File {
				if fd < 0 {
					return nil
				}
				return &File{fd, name}
			}
			// 调用
			f := NewFile(10, "./test.txt")

	* 查看结构体类型T的一个实例占用了多少内存
		size := unsafe.Sizeof(T{})


*/

