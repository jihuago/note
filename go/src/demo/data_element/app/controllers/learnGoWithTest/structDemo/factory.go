package structDemo

type File struct {
	fd int
	name string // 文件名
}

// File类型对应的工厂方法
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}

// 如何强制使用工厂方法
// 1. 私有化结构体
// 2. 提供一个工厂方法给其他包使用
type matrix struct {

}

func NewMatrix(params interface{}) *matrix  {
	m := new(matrix)
	return m
}

// Unix 网络编程卷一  => 深入理解 Nginx  => Redis 设计与实现
// go基础知识  => Gin源码 => GPRC

