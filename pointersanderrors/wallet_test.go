package main

import "testing"


func TestWallet(t *testing.T) { 
  t.Run("deposit", func(t *testing.T) {
    wallet := Wallet{}  
    wallet.Deposit(Bitcoin(10))
 
    assertBalance(t, wallet, Bitcoin(10))
  }) 
  t.Run("with draw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(12)}
    err := wallet.WithDraw(Bitcoin(4))
    
    assertNoError(t, err)
    assertBalance(t, wallet, Bitcoin(8))
  }) 
  
  t.Run("withdraw insufficient funds", func(t *testing.T) {
    startBalance := Bitcoin(20)

    wallet := Wallet{startBalance}
    err := wallet.WithDraw(Bitcoin(100))
    
    assertError(t, err, ErrInsufficientFunds)  
    assertBalance(t, wallet, startBalance)
  })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
      t.Errorf("Expected %s but got %s", want, got)
    }
  }

func assertError(t testing.TB, got, want error) {
    t.Helper()
    if got == nil {
      t.Fatal("wanted an error but didn't get one")
    }

    if got != want {
      t.Errorf("Got error '%s' but wanted '%s'", got, want)
    }
}

func assertNoError(t testing.TB, err error) {
  t.Helper()
  if err != nil {
    t.Errorf("Wanted no error but got %s", err.Error())
  }
}


