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
