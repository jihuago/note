package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T)  {

	t.Run("sleep after every print", func(t *testing.T) {
		spySleeper := &CountdownOperationsSpy{}
		Countdown(spySleeper, spySleeper)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
		}

	})

	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		// 更新测试以注入对我们监视器的依赖，并断言sleep被调用了4次
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

}
