package common

func GetStringFirstChar(str string) (string, int) {
	// 将string转为字符类型数组
	runStr := []rune(str)

	return string(runStr[0]), len(runStr)
}



