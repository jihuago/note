package goroutine

import "testing"

func TestQuery(t *testing.T) {
	t.Run("test multi query", func(t *testing.T) {
		Query()
	})
}
