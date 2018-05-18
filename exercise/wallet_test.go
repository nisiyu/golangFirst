package wallet

import "testing"

func WalletTest(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("want %v but get %v", want, got)
		}
	})

	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}	
		wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("want %v but get %v", want, got)
		}

	})

	t.Run("WithDraw error", func(t *testing.T) {

		t.Errorf("aaa")
		// wallet := Wallet{balance: Bitcoin(20)}	
		// err := wallet.Withdraw(30)

		// got := wallet.Balance()
		// want := Bitcoin(20)

		// if got != want {
		// 	t.Errorf("want %v but get %v", want, got)
		// }

		// if err != nil {
		// 	t.Errorf("exception excepted")
		// }

	})
	
}