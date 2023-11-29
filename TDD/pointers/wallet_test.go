package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("BitcoinWallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
