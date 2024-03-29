package main

import "fmt" 

func (w *Wallet) Deposit(amount Bitcoin) { 
  w.balance += amount 
}

func (w Wallet) Balance() Bitcoin {
  return w.balance
}

func (w *Wallet) WithDraw(amount Bitcoin) error {
  if w.balance < amount {
    return ErrInsufficientFunds
  }
  w.balance -= amount
  return nil
}

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}

