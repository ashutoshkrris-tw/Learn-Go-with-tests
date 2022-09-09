package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(100))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(500)}

		wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(400))
	})
}

func assertBalance(t testing.TB, wallet Wallet, wanted Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != wanted {
		t.Errorf("\nActual: %s\nExpected: %s", got, wanted)
	}
}
