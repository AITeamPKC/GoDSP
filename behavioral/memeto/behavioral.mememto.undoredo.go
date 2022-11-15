package main

import "fmt"

type Memeto struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memeto
	current int
}

func (b *BankAccount) String() string {
	return fmt.Sprint("Balance = $", b.balance, ", current = ", b.current)
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{balance: balance}
	b.changes = append(b.changes, &Memeto{balance})
	return b
}

func (b *BankAccount) Deposit(amount int) *Memeto {
	b.balance += amount
	m := Memeto{b.balance}
	b.changes = append(b.changes, &m)
	b.current++
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
	return &m
}

func (b *BankAccount) Restore(m *Memeto) {
	if m != nil {
		b.balance -= m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Memeto {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Memeto {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

//Tao 1 slice de luu bo nho => undo redo tuong ung tien lui

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 2:", ba)
	ba.Redo()
	fmt.Println("Redo:", ba)
}
