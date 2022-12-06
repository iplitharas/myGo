package PointersErrors

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s - want %s", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{100}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(90))

	})
	t.Run("withdraw insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}
