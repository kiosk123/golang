package banking

import (
	"errors"
	"fmt"
)

// account struct
type account struct {
	owner   string
	balance int
}

var erroNoMoney = errors.New("Can't withdraw you are poor")

func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a *account) Balance() int {
	return a.balance
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return erroNoMoney
	} else {
		a.balance -= amount
	}
	return nil
}

func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.owner, " has ", a.balance)
}
