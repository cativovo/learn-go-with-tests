package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(100),
		}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoErr(t, err)
		assertBalance(t, wallet, Bitcoin(90))
	})

	t.Run("withdraw insufficient balance", func(t *testing.T) {
		startingBalance := 20
		wallet := Wallet{balance: Bitcoin(startingBalance)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(startingBalance))
		assertErr(t, err, ErrInsufficientBalance)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoErr(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Error("got an error but didn't want one")
	}
}

func assertErr(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Error("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
