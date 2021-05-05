package goroutine

import "testing"

func TestRunTimer(t *testing.T) {
/*	t.Run("test ticker", func(t *testing.T) {
		RunTimer()
	})*/
	
	t.Run("test ticker update", func(t *testing.T) {
		LimitRequest()
	})
}
