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
