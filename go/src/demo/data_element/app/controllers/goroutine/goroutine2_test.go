package goroutine

import "testing"

func TestRunGoroutine2(t *testing.T) {
	t.Run("test goroutine2", func(t *testing.T) {
		RunGoroutine2()
	})
	
	t.Run("test block", func(t *testing.T) {
		Get()
	})

	t.Run("test block2", func(t *testing.T) {
		RunChannelBlock3()
	})

	t.Run("blocking", func(t *testing.T) {
		RunF1()
	})

	t.Run("sumChannel", func(t *testing.T) {
		DoSum()
	})
}
