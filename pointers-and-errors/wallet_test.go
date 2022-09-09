package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(100))
	})

	t.Run("should withdraw when sufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(500)}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(400))
		assertNoError(t, err)
	})

	t.Run("should not withdraw when insufficient balance", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(500))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, wanted Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != wanted {
		t.Errorf("\nActual: %s\nExpected: %s", got, wanted)
	}
}

func assertError(t testing.TB, got error, wanted error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != wanted {
		t.Errorf("\nActual: %q\nExpected: %q", got, wanted)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
