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
	var actual float64 = 10

	//Act
	account.Deposit(actual)

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
