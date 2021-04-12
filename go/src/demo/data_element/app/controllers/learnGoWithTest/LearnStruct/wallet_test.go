package LearnStruct

import "testing"

func TestWallet(t *testing.T)  {

	// 为错误检查做一个快速测试的助手方法，帮助我们的测试读起来更清晰
	assertError := func(t *testing.T, err error) {
		if err == nil {
			t.Error("wanted an error but didnt get one")
		}
	}

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(20)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err)

	})
}
