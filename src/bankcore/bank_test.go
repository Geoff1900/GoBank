package bank

import (
	"strings"
	"testing"
)

func TestBankCore(t *testing.T) {
	expected := "TEST RESPONSE"
	actual := strings.ToUpper(Ping())
	if actual != expected {
		t.Errorf("Ping() returned %s when %s was expected", actual, expected)
	}
}

func TestAccount(t *testing.T) {
	//Arrange
	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	expected := "GEOFF"
	//Act
	actual := strings.ToUpper(account.Name)
	//Assert
	if actual != expected {
		t.Errorf("Account() returned %s when %s was expected", actual, expected)
	}

}

func TestDeposit(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	var expected float64 = 10

	//Act
	account.Deposit(10.0)
	var actual float64 = account.Balance

	//Assert
	if actual != expected {
		t.Errorf("Account() returned %f when %f was expected", actual, expected)
	}
}

func TestDepositLessThanZero(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	var actual float64 = -10

	//Act

	//Assert
	if err := account.Deposit(actual); err == nil {
		t.Errorf("Account() error expected if deposit is less than zero")
	}
}

func TestDepositEqualsZero(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	var actual float64 = 0

	//Act

	//Assert
	if err := account.Deposit(actual); err == nil {
		t.Errorf("Account() error expected if deposit is equal to zero")
	}
}

func TestWithdraw(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	var expected float64 = 0.0

	//Act
	account.Withdraw(10.0)
	var actual float64 = account.Balance

	//Assert
	if actual != expected {
		t.Errorf("Account() returned %f when %f was expected", actual, expected)
	}
}

func TestWithdrawLessThanZero(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	var actual float64 = -10

	//Act

	//Assert
	if account.Withdraw(actual) == nil {
		t.Errorf("Account() error expected if withdrawal is less than zero")
	}
}

func TestWithdrawEqualsZero(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	var actual float64 = 0

	//Act

	//Assert
	if err := account.Withdraw(actual); err == nil {
		t.Errorf("Account() error expected if withdrawal is equal to zero")
	}
}

func TestWithdrawLessThanBalance(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "123",
		},
		Number:  001,
		Balance: 0,
	}
	var actual float64 = 10

	//Act

	//Assert
	if err := account.Withdraw(actual); err == nil {
		t.Errorf("Account() error expected if withdrawal is greater than balance")
	}
}

func TestStatement(t *testing.T) {
	//Arrange

	account := Account{
		Customer: Customer{
			Name:    "Geoff",
			Address: "Porthcawl",
			Phone:   "12345",
		},
		Number:  01,
		Balance: 0.0,
	}
	account.Deposit(1.0)
	//Act
	actual := account.Statement()
	expected := "1 - Geoff - 1"

	//Assert
	if actual != expected {
		t.Errorf("Account. Statement() returned %s when %s was expected", actual, expected)
	}
}
