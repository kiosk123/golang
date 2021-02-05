package main

import (
	"fmt"

	"github.com/kiosk123/golang/banking"
)

func main() {
	account := banking.NewAccount("nico")
	fmt.Println(account)

	account.Deposit(100)
	fmt.Println(account)

	fmt.Println(account.Balance())
	err := account.Withdraw(200)

	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
}
