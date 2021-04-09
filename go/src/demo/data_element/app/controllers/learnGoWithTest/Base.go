package learnGoWithTest

// 通常定义一个常量，常量可以提高应用程序的性能，避免了每次使用Hello()时创建"Hello, "字符串示例
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

