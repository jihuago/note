package learnGoWithTest


func Repeat(character string, repeatTime int) string  {
	var expected string

	for i := 0; i < repeatTime; i++ {
		expected += character
	}

	return expected
}
