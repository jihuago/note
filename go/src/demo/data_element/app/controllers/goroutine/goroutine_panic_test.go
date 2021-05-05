package goroutine

import "testing"

// go test -run TestRunTel
func TestRunTel(t *testing.T) {
/*	t.Run("test Runtel function", func(t *testing.T) {
		RunTel()
	})*/
	
	t.Run("test RandomBitgen", func(t *testing.T) {
		RandomBitgen()
	})
}
