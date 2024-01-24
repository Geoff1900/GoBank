package bank

import (
	"errors"
)

func Ping() string {
	return "TEST RESPONSE"
}

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (account *Account) Deposit(deposit float64) error {
	if deposit <= 0 {
		return errors.New("value of desposit must be greater than 0")
	}
	account.Balance += deposit
	return nil
}
