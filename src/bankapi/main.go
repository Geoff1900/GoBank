package main

import (
	"fmt"

	bank "github.com/Geoff1900/GoBank/src/bankcore"
)

func main() {
	fmt.Println(bank.Ping())
	var accounts = map[float64]*bank.Account{}

	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

}
