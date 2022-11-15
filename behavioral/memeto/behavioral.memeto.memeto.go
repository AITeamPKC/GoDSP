package main

import "fmt"

type Memeto struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) *Memeto {
	b.balance += amount
	return &Memeto{b.balance}
}

func (b *BankAccount) Restore(m *Memeto) {
	b.balance = m.Balance
}

func main() {
	ba := BankAccount{100}
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)

	fmt.Println(ba)

	ba.Restore(m1)
	fmt.Println(ba)

	ba.Restore(m2)
	fmt.Println(ba)
}
