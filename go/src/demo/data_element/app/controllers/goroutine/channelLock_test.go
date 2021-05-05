package goroutine

import "testing"

func TestRunLock(t *testing.T) {
	t.Run("test lock", func(t *testing.T) {
		RunLock()
	})
}
