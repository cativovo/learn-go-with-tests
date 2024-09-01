package revisitingarraysandsliceswithgenerics

import (
	"fmt"
	"testing"
)

func TestBadBank(t *testing.T) {
	john := Account{
		Name:    "John",
		Balance: 100,
	}
	juan := Account{
		Name:    "Juan",
		Balance: 75,
	}
	jose := Account{
		Name:    "Jose",
		Balance: 200,
	}

	transactions := []Transaction{
		NewTransaction(juan, john, 100),
		NewTransaction(jose, juan, 25),
	}

	testCases := []struct {
		account Account
		want    float64
	}{
		{
			account: john,
			want:    200,
		},
		{
			account: juan,
			want:    0,
		},
		{
			account: jose,
			want:    175,
		},
	}

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := newBalanceFor(testCase.account)
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
