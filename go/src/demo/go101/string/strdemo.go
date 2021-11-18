package string1

import (
	"fmt"
	"strings"
)

func Stringdemo()  {
	var builder strings.Builder

	builder.WriteString("Hello world!")
	fmt.Println(builder.String())
}
