package lock

import "testing"

func TestWithdraw(t *testing.T) {

	t.Run("test one", func(t *testing.T) {
		want := false
		got := Withdraw(0)

		if got != want {
			t.Errorf("got '%t', want '%t'", got, want)
		}
	})

}
