package main

import (
	"fmt"

	"github.com/kiosk123/golang/banking"
	"github.com/kiosk123/golang/dict"
)

func main() {

	// -- struct --
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
	fmt.Println(account)

	// -- struct --
	// -- Dictionary --

	dictionary := dict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)

	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(definition)
	} else {
		fmt.Println(err)
	}
	// -- Dictionary --
}
