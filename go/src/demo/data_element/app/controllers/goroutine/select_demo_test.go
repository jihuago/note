package goroutine

import "testing"

func TestRunSelectDemo(t *testing.T) {
	t.Run("test select", func(t *testing.T) {
		RunSelectDemo()
	})
}
