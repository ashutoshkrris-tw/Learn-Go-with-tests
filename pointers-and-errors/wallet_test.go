package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(100)

	got := wallet.Balance()
	wanted := Bitcoin(100)

	if got != wanted {
		t.Errorf("\nActual: %s\nExpected: %s", got, wanted)
	}
}
