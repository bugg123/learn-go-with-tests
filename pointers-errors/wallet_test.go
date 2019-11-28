package wallet

import (
	"testing"
)

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Error("didn't get an error and expected one")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error("got an error and didn't expected one")
	}
}

func assertBalance(t *testing.T, wallet Wallet, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("%#v expected %s but got %s", wallet, want, got)
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, wallet, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, wallet, got, want)
		assertNoError(t, err)
	})

	t.Run("withdraw over balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(30)

		got := wallet.Balance()
		want := Bitcoin(20)

		assertBalance(t, wallet, got, want)
		assertError(t, err, ErrInsufficientFunds)
	})

}
