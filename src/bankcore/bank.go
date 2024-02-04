package bank

import (
	"errors"
	"fmt"
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

func (account *Account) Withdraw(withdrawal float64) error {
	if withdrawal <= 0 {
		return errors.New("value of withdrawal must be greater than 0")
	}
	if withdrawal > account.Balance {
		return errors.New("value of withdrawal must be less than balance")
	}
	account.Balance -= withdrawal
	return nil
}

func (account *Account) Statement() string {
	statement := fmt.Sprintf("%d - %s - %v", account.Number, account.Name, account.Balance)
	return statement
}

// Mock Transfer method
func (accountTo *Account) Transfer(amount float64, accountFrom *Account) error {
	// Do nothing
	if err := accountTo.Withdraw(amount); err != nil {
		return errors.New("value of withdrawal must be less than balance")
	}

	accountFrom.Deposit(amount)
	return nil
}
