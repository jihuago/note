package _interface

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {

}

func (s *Student) Speak(think string) (talk string)  {
	if think == "sb" {
		talk = "ni shi da shui bi"
	} else {
		talk = "你好"
	}

	return
}

func Init3()  {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
