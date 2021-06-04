package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var noMoney = errors.New("Can't withdraw")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return noMoney
	}
	a.balance -= amount
	return nil
}

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of account
func (a Account) Owner() string {
	return a.owner
}

// Check balance
func (a Account) Balance() int {
	return a.balance
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account\nHas: ", a.Balance())
}
